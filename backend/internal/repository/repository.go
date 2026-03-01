package repository

import (
	"gorm.io/gorm"

	"aiwriter/internal/model"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

type WorkRepository struct {
	db *gorm.DB
}

func NewWorkRepository(db *gorm.DB) *WorkRepository {
	return &WorkRepository{db: db}
}

func (r *WorkRepository) Create(work *model.Work) error {
	return r.db.Create(work).Error
}

func (r *WorkRepository) FindByID(id uint) (*model.Work, error) {
	var work model.Work
	err := r.db.First(&work, id).Error
	return &work, err
}

func (r *WorkRepository) FindByUserID(userID uint, page, pageSize int) ([]model.Work, int64, error) {
	var works []model.Work
	var total int64

	query := r.db.Model(&model.Work{}).Where("user_id = ? AND deleted_at IS NULL", userID)
	query.Count(&total)

	err := query.Offset((page - 1) * pageSize).Limit(pageSize).Order("updated_at DESC").Find(&works).Error
	return works, total, err
}

func (r *WorkRepository) Update(work *model.Work) error {
	return r.db.Save(work).Error
}

func (r *WorkRepository) Delete(id uint) error {
	return r.db.Delete(&model.Work{}, id).Error
}

type VolumeRepository struct {
	db *gorm.DB
}

func NewVolumeRepository(db *gorm.DB) *VolumeRepository {
	return &VolumeRepository{db: db}
}

func (r *VolumeRepository) Create(volume *model.Volume) error {
	return r.db.Create(volume).Error
}

func (r *VolumeRepository) FindByID(id uint) (*model.Volume, error) {
	var volume model.Volume
	err := r.db.First(&volume, id).Error
	return &volume, err
}

func (r *VolumeRepository) FindByWorkID(workID uint) ([]model.Volume, error) {
	var volumes []model.Volume
	err := r.db.Where("work_id = ? AND deleted_at IS NULL", workID).Order("volume_index ASC").Find(&volumes).Error
	return volumes, err
}

func (r *VolumeRepository) Update(volume *model.Volume) error {
	return r.db.Save(volume).Error
}

func (r *VolumeRepository) Delete(id uint) error {
	return r.db.Delete(&model.Volume{}, id).Error
}

type ChapterRepository struct {
	db *gorm.DB
}

func NewChapterRepository(db *gorm.DB) *ChapterRepository {
	return &ChapterRepository{db: db}
}

func (r *ChapterRepository) Create(chapter *model.Chapter) error {
	return r.db.Create(chapter).Error
}

func (r *ChapterRepository) FindByID(id uint) (*model.Chapter, error) {
	var chapter model.Chapter
	err := r.db.First(&chapter, id).Error
	return &chapter, err
}

func (r *ChapterRepository) FindByWorkID(workID uint) ([]model.Chapter, error) {
	var chapters []model.Chapter
	err := r.db.Where("work_id = ? AND deleted_at IS NULL", workID).Order("chapter_index ASC").Find(&chapters).Error
	return chapters, err
}

func (r *ChapterRepository) FindByVolumeID(volumeID uint) ([]model.Chapter, error) {
	var chapters []model.Chapter
	err := r.db.Where("volume_id = ? AND deleted_at IS NULL", volumeID).Order("chapter_index ASC").Find(&chapters).Error
	return chapters, err
}

func (r *ChapterRepository) Update(chapter *model.Chapter) error {
	return r.db.Save(chapter).Error
}

func (r *ChapterRepository) Delete(id uint) error {
	return r.db.Delete(&model.Chapter{}, id).Error
}

type Repositories struct {
	User    *UserRepository
	Work    *WorkRepository
	Volume  *VolumeRepository
	Chapter *ChapterRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:    NewUserRepository(db),
		Work:    NewWorkRepository(db),
		Volume:  NewVolumeRepository(db),
		Chapter: NewChapterRepository(db),
	}
}
