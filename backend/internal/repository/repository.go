package repository

import (
	"context"
	"fmt"

	"aiwriter/internal/model"
	"aiwriter/internal/storage"
)

type Repository struct {
	store   storage.Storage
	meta    *model.MetaStore
	metaKey string
}

func NewRepository(store storage.Storage) *Repository {
	return &Repository{
		store:   store,
		meta:    model.NewMetaStore(),
		metaKey: "metadata/store.json",
	}
}

func (r *Repository) LoadMeta(ctx context.Context) error {
	err := r.store.GetMeta(ctx, r.metaKey, r.meta)
	if err != nil {
		r.meta = model.NewMetaStore()
	}
	return nil
}

func (r *Repository) SaveMeta(ctx context.Context) error {
	return r.store.PutMeta(ctx, r.metaKey, r.meta)
}

func (r *Repository) getNextID(key string) uint {
	id := r.meta.NextIDs[key]
	r.meta.NextIDs[key] = id + 1
	return id
}

type UserRepository struct {
	repo *Repository
}

func NewUserRepository(repo *Repository) *UserRepository {
	return &UserRepository{repo: repo}
}

func (r *UserRepository) Create(user *model.User) error {
	ctx := context.Background()
	user.ID = r.repo.getNextID("user")
	user.CreatedAt = user.CreatedAt
	user.UpdatedAt = user.UpdatedAt
	r.repo.meta.Users[fmt.Sprintf("%d", user.ID)] = *user
	return r.repo.SaveMeta(ctx)
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	user, ok := r.repo.meta.Users[fmt.Sprintf("%d", id)]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	for _, user := range r.repo.meta.Users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, user := range r.repo.meta.Users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (r *UserRepository) Update(user *model.User) error {
	ctx := context.Background()
	r.repo.meta.Users[fmt.Sprintf("%d", user.ID)] = *user
	return r.repo.SaveMeta(ctx)
}

type WorkRepository struct {
	repo *Repository
}

func NewWorkRepository(repo *Repository) *WorkRepository {
	return &WorkRepository{repo: repo}
}

func (r *WorkRepository) Create(work *model.Work) error {
	ctx := context.Background()
	work.ID = r.repo.getNextID("work")
	work.CreatedAt = work.CreatedAt
	work.UpdatedAt = work.UpdatedAt
	r.repo.meta.Works[fmt.Sprintf("%d", work.ID)] = *work
	return r.repo.SaveMeta(ctx)
}

func (r *WorkRepository) FindByID(id uint) (*model.Work, error) {
	work, ok := r.repo.meta.Works[fmt.Sprintf("%d", id)]
	if !ok {
		return nil, fmt.Errorf("work not found")
	}
	return &work, nil
}

func (r *WorkRepository) FindByUserID(userID uint, page, pageSize int) ([]model.Work, int64, error) {
	var works []model.Work
	var total int64
	for _, work := range r.repo.meta.Works {
		if work.UserID == userID {
			works = append(works, work)
			total++
		}
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if start > len(works) {
		works = []model.Work{}
	} else {
		if end > len(works) {
			end = len(works)
		}
		works = works[start:end]
	}
	return works, total, nil
}

func (r *WorkRepository) Update(work *model.Work) error {
	ctx := context.Background()
	r.repo.meta.Works[fmt.Sprintf("%d", work.ID)] = *work
	return r.repo.SaveMeta(ctx)
}

func (r *WorkRepository) Delete(id uint) error {
	ctx := context.Background()
	delete(r.repo.meta.Works, fmt.Sprintf("%d", id))
	return r.repo.SaveMeta(ctx)
}

type VolumeRepository struct {
	repo *Repository
}

func NewVolumeRepository(repo *Repository) *VolumeRepository {
	return &VolumeRepository{repo: repo}
}

func (r *VolumeRepository) Create(volume *model.Volume) error {
	ctx := context.Background()
	volume.ID = r.repo.getNextID("volume")
	volume.CreatedAt = volume.CreatedAt
	volume.UpdatedAt = volume.UpdatedAt
	r.repo.meta.Volumes[fmt.Sprintf("%d", volume.ID)] = *volume
	return r.repo.SaveMeta(ctx)
}

func (r *VolumeRepository) FindByID(id uint) (*model.Volume, error) {
	volume, ok := r.repo.meta.Volumes[fmt.Sprintf("%d", id)]
	if !ok {
		return nil, fmt.Errorf("volume not found")
	}
	return &volume, nil
}

func (r *VolumeRepository) FindByWorkID(workID uint) ([]model.Volume, error) {
	var volumes []model.Volume
	for _, volume := range r.repo.meta.Volumes {
		if volume.WorkID == workID {
			volumes = append(volumes, volume)
		}
	}
	return volumes, nil
}

func (r *VolumeRepository) Update(volume *model.Volume) error {
	ctx := context.Background()
	r.repo.meta.Volumes[fmt.Sprintf("%d", volume.ID)] = *volume
	return r.repo.SaveMeta(ctx)
}

func (r *VolumeRepository) Delete(id uint) error {
	ctx := context.Background()
	delete(r.repo.meta.Volumes, fmt.Sprintf("%d", id))
	return r.repo.SaveMeta(ctx)
}

type ChapterRepository struct {
	repo *Repository
}

func NewChapterRepository(repo *Repository) *ChapterRepository {
	return &ChapterRepository{repo: repo}
}

func (r *ChapterRepository) Create(chapter *model.Chapter) error {
	ctx := context.Background()
	chapter.ID = r.repo.getNextID("chapter")
	chapter.CreatedAt = chapter.CreatedAt
	chapter.UpdatedAt = chapter.UpdatedAt
	r.repo.meta.Chapters[fmt.Sprintf("%d", chapter.ID)] = *chapter
	return r.repo.SaveMeta(ctx)
}

func (r *ChapterRepository) FindByID(id uint) (*model.Chapter, error) {
	chapter, ok := r.repo.meta.Chapters[fmt.Sprintf("%d", id)]
	if !ok {
		return nil, fmt.Errorf("chapter not found")
	}
	return &chapter, nil
}

func (r *ChapterRepository) FindByWorkID(workID uint) ([]model.Chapter, error) {
	var chapters []model.Chapter
	for _, chapter := range r.repo.meta.Chapters {
		if chapter.WorkID == workID {
			chapters = append(chapters, chapter)
		}
	}
	return chapters, nil
}

func (r *ChapterRepository) FindByVolumeID(volumeID uint) ([]model.Chapter, error) {
	var chapters []model.Chapter
	for _, chapter := range r.repo.meta.Chapters {
		if chapter.VolumeID == volumeID {
			chapters = append(chapters, chapter)
		}
	}
	return chapters, nil
}

func (r *ChapterRepository) Update(chapter *model.Chapter) error {
	ctx := context.Background()
	r.repo.meta.Chapters[fmt.Sprintf("%d", chapter.ID)] = *chapter
	return r.repo.SaveMeta(ctx)
}

func (r *ChapterRepository) Delete(id uint) error {
	ctx := context.Background()
	delete(r.repo.meta.Chapters, fmt.Sprintf("%d", id))
	return r.repo.SaveMeta(ctx)
}

type Repositories struct {
	User    *UserRepository
	Work    *WorkRepository
	Volume  *VolumeRepository
	Chapter *ChapterRepository
}

func NewRepositories(store storage.Storage) *Repositories {
	repo := NewRepository(store)
	repo.LoadMeta(context.Background())
	return &Repositories{
		User:    NewUserRepository(repo),
		Work:    NewWorkRepository(repo),
		Volume:  NewVolumeRepository(repo),
		Chapter: NewChapterRepository(repo),
	}
}
