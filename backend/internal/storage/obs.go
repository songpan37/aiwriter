package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

type Storage interface {
	PutObject(ctx context.Context, key string, data []byte) error
	GetObject(ctx context.Context, key string) ([]byte, error)
	DeleteObject(ctx context.Context, key string) error
	ListObjects(ctx context.Context, prefix string) ([]string, error)
	ObjectExists(ctx context.Context, key string) (bool, error)
	GetMeta(ctx context.Context, key string, v interface{}) error
	PutMeta(ctx context.Context, key string, v interface{}) error
	DeleteMeta(ctx context.Context, key string, v interface{}) error
}

type OBSClient struct {
	client *obs.ObsClient
	bucket string
	mu     sync.RWMutex
	cache  map[string]interface{}
}

type LocalStorage struct {
	baseDir string
	mu      sync.RWMutex
}

func NewLocalStorage(baseDir string) *LocalStorage {
	os.MkdirAll(baseDir, 0755)
	return &LocalStorage{baseDir: baseDir}
}

func (s *LocalStorage) getPath(key string) string {
	return filepath.Join(s.baseDir, key)
}

func (s *LocalStorage) PutObject(ctx context.Context, key string, data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	path := s.getPath(key)
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func (s *LocalStorage) GetObject(ctx context.Context, key string) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return os.ReadFile(s.getPath(key))
}

func (s *LocalStorage) DeleteObject(ctx context.Context, key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return os.Remove(s.getPath(key))
}

func (s *LocalStorage) ListObjects(ctx context.Context, prefix string) ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var keys []string
	_ = s.getPath(prefix)
	err := filepath.Walk(s.baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			rel, err := filepath.Rel(s.baseDir, path)
			if err != nil {
				return err
			}
			keys = append(keys, rel)
		}
		return nil
	})
	return keys, err
}

func (s *LocalStorage) ObjectExists(ctx context.Context, key string) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, err := os.Stat(s.getPath(key))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *LocalStorage) GetMeta(ctx context.Context, key string, v interface{}) error {
	data, err := s.GetObject(ctx, key)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (s *LocalStorage) PutMeta(ctx context.Context, key string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return s.PutObject(ctx, key, data)
}

func (s *LocalStorage) DeleteMeta(ctx context.Context, key string, v interface{}) error {
	return s.DeleteObject(ctx, key)
}

func NewOBSClient(endpoint, region, bucket, accessKey, secretKey string) (*OBSClient, error) {
	obsClient, err := obs.New(accessKey, secretKey, endpoint, obs.WithSignature(obs.SignatureObs))
	if err != nil {
		return nil, fmt.Errorf("failed to create OBS client: %w", err)
	}

	createBucketInput := &obs.CreateBucketInput{}
	createBucketInput.Bucket = bucket
	createBucketInput.Location = region
	_, err = obsClient.CreateBucket(createBucketInput)
	if err != nil {
		errMsg := err.Error()
		if !strings.Contains(errMsg, "NoSuchBucket") && !strings.Contains(errMsg, "BucketAlreadyExists") {
			return nil, fmt.Errorf("failed to create bucket: %w", err)
		}
	}

	return &OBSClient{
		client: obsClient,
		bucket: bucket,
		cache:  make(map[string]interface{}),
	}, nil
}

func (s *OBSClient) PutObject(ctx context.Context, key string, data []byte) error {
	input := &obs.PutObjectInput{}
	input.Bucket = s.bucket
	input.Key = key
	input.Body = bytes.NewReader(data)

	_, err := s.client.PutObject(input)
	if err != nil {
		return fmt.Errorf("failed to put object: %w", err)
	}
	return nil
}

func (s *OBSClient) GetObject(ctx context.Context, key string) ([]byte, error) {
	input := &obs.GetObjectInput{}
	input.Bucket = s.bucket
	input.Key = key

	output, err := s.client.GetObject(input)
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %w", err)
	}
	defer output.Body.Close()

	data, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read object body: %w", err)
	}
	return data, nil
}

func (s *OBSClient) DeleteObject(ctx context.Context, key string) error {
	input := &obs.DeleteObjectInput{}
	input.Bucket = s.bucket
	input.Key = key

	_, err := s.client.DeleteObject(input)
	if err != nil {
		return fmt.Errorf("failed to delete object: %w", err)
	}
	return nil
}

func (s *OBSClient) ListObjects(ctx context.Context, prefix string) ([]string, error) {
	input := &obs.ListObjectsInput{}
	input.Bucket = s.bucket
	input.Prefix = prefix

	output, err := s.client.ListObjects(input)
	if err != nil {
		return nil, fmt.Errorf("failed to list objects: %w", err)
	}

	keys := make([]string, 0, len(output.Contents))
	for _, obj := range output.Contents {
		keys = append(keys, obj.Key)
	}
	return keys, nil
}

func (s *OBSClient) ObjectExists(ctx context.Context, key string) (bool, error) {
	input := &obs.HeadObjectInput{}
	input.Bucket = s.bucket
	input.Key = key

	_, err := s.client.HeadObject(input)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (s *OBSClient) GetMeta(ctx context.Context, key string, v interface{}) error {
	data, err := s.GetObject(ctx, key)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (s *OBSClient) PutMeta(ctx context.Context, key string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return s.PutObject(ctx, key, data)
}

func (s *OBSClient) DeleteMeta(ctx context.Context, key string, v interface{}) error {
	return s.DeleteObject(ctx, key)
}

func EnsureDataDir(sqlitePath string) error {
	dir := sqlitePath[:len(sqlitePath)-len("/aiwriter.db")]
	if dir == sqlitePath {
		dir = "."
	}
	return os.MkdirAll(dir, 0755)
}
