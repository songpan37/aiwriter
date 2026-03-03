package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage interface {
	PutObject(ctx context.Context, key string, data []byte) error
	GetObject(ctx context.Context, key string) ([]byte, error)
	DeleteObject(ctx context.Context, key string) error
	ListObjects(ctx context.Context, prefix string) ([]string, error)
	ObjectExists(ctx context.Context, key string) (bool, error)
	GetMeta(ctx context.Context, key string, v interface{}) error
	PutMeta(ctx context.Context, key string, v interface{}) error
	DeleteMeta(ctx context.Context, key string) error
}

type S3Client struct {
	client *s3.Client
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

func (s *LocalStorage) DeleteMeta(ctx context.Context, key string) error {
	return s.DeleteObject(ctx, key)
}

func NewS3Client(endpoint, region, bucket, accessKey, secretKey string) (*S3Client, error) {
	awsEndpoint := fmt.Sprintf("https://%s", endpoint)

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               awsEndpoint,
			HostnameImmutable: true,
		}, nil
	})

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return &S3Client{
		client: client,
		bucket: bucket,
		cache:  make(map[string]interface{}),
	}, nil
}

func (s *S3Client) PutObject(ctx context.Context, key string, data []byte) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})
	return err
}

func (s *S3Client) GetObject(ctx context.Context, key string) ([]byte, error) {
	result, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()
	return io.ReadAll(result.Body)
}

func (s *S3Client) DeleteObject(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	return err
}

func (s *S3Client) ListObjects(ctx context.Context, prefix string) ([]string, error) {
	result, err := s.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(result.Contents))
	for _, obj := range result.Contents {
		if obj.Key != nil {
			keys = append(keys, *obj.Key)
		}
	}
	return keys, nil
}

func (s *S3Client) ObjectExists(ctx context.Context, key string) (bool, error) {
	_, err := s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (s *S3Client) GetMeta(ctx context.Context, key string, v interface{}) error {
	data, err := s.GetObject(ctx, key)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (s *S3Client) PutMeta(ctx context.Context, key string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return s.PutObject(ctx, key, data)
}

func (s *S3Client) DeleteMeta(ctx context.Context, key string) error {
	return s.DeleteObject(ctx, key)
}

func EnsureDataDir(sqlitePath string) error {
	dir := sqlitePath[:len(sqlitePath)-len("/aiwriter.db")]
	if dir == sqlitePath {
		dir = "."
	}
	return os.MkdirAll(dir, 0755)
}
