package service

import (
	"errors"

	"aiwriter/internal/config"
	"aiwriter/internal/dto/request"
	"aiwriter/internal/dto/response"
	"aiwriter/internal/model"
	"aiwriter/internal/repository"
	"aiwriter/pkg/utils"

	"gorm.io/gorm"
)

type Service struct {
	repos  *repository.Repositories
	config *config.Config
}

func NewServices(repos *repository.Repositories, cfg *config.Config) *Service {
	return &Service{
		repos:  repos,
		config: cfg,
	}
}

type AuthService struct {
	repos  *repository.Repositories
	config *config.Config
}

func NewAuthService(repos *repository.Repositories, cfg *config.Config) *AuthService {
	return &AuthService{
		repos:  repos,
		config: cfg,
	}
}

func (s *AuthService) Register(req *request.RegisterRequest) (*response.LoginResponse, error) {
	existingUser, _ := s.repos.User.FindByUsername(req.Username)
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	existingEmail, _ := s.repos.User.FindByEmail(req.Email)
	if existingEmail != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	if err := s.repos.User.Create(user); err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWT(user.ID, user.Username, s.config.JWTSecret)
	if err != nil {
		return nil, err
	}

	return &response.LoginResponse{
		Token: token,
		User: response.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Avatar:   user.Avatar,
		},
	}, nil
}

func (s *AuthService) Login(req *request.LoginRequest) (*response.LoginResponse, error) {
	user, err := s.repos.User.FindByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username or password")
		}
		return nil, err
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Username, s.config.JWTSecret)
	if err != nil {
		return nil, err
	}

	return &response.LoginResponse{
		Token: token,
		User: response.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Avatar:   user.Avatar,
		},
	}, nil
}

type WorkService struct {
	repos *repository.Repositories
}

func NewWorkService(repos *repository.Repositories) *WorkService {
	return &WorkService{repos: repos}
}

func (s *WorkService) Create(userID uint, req *request.CreateWorkRequest) (*model.Work, error) {
	work := &model.Work{
		UserID:     userID,
		CategoryID: req.CategoryID,
		Title:      req.Title,
		Cover:      req.Cover,
		Intro:      req.Intro,
	}

	if err := s.repos.Work.Create(work); err != nil {
		return nil, err
	}

	return work, nil
}

func (s *WorkService) GetList(userID uint, page, pageSize int) (*response.WorkListResponse, error) {
	works, total, err := s.repos.Work.FindByUserID(userID, page, pageSize)
	if err != nil {
		return nil, err
	}

	var workResponses []response.WorkResponse
	for _, work := range works {
		workResponses = append(workResponses, response.WorkResponse{
			ID:           work.ID,
			Title:        work.Title,
			Cover:        work.Cover,
			ChapterCount: work.ChapterCount,
			WordCount:    work.WordCount,
			UpdatedAt:    work.UpdatedAt,
		})
	}

	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	return &response.WorkListResponse{
		List: workResponses,
		Pagination: response.Pagination{
			Page:       page,
			PageSize:   pageSize,
			Total:      int(total),
			TotalPages: totalPages,
		},
	}, nil
}

func (s *WorkService) GetByID(id uint) (*model.Work, error) {
	return s.repos.Work.FindByID(id)
}

func (s *WorkService) Update(id uint, req *request.UpdateWorkRequest) (*model.Work, error) {
	work, err := s.repos.Work.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		work.Title = req.Title
	}
	if req.Cover != "" {
		work.Cover = req.Cover
	}
	if req.Intro != "" {
		work.Intro = req.Intro
	}

	if err := s.repos.Work.Update(work); err != nil {
		return nil, err
	}

	return work, nil
}

func (s *WorkService) Delete(id uint) error {
	return s.repos.Work.Delete(id)
}

type VolumeService struct {
	repos *repository.Repositories
}

func NewVolumeService(repos *repository.Repositories) *VolumeService {
	return &VolumeService{repos: repos}
}

func (s *VolumeService) Create(workID uint, req *request.CreateVolumeRequest) (*model.Volume, error) {
	volume := &model.Volume{
		WorkID:      workID,
		VolumeIndex: req.VolumeIndex,
		Title:       req.Title,
		Summary:     req.Summary,
	}

	if err := s.repos.Volume.Create(volume); err != nil {
		return nil, err
	}

	return volume, nil
}

func (s *VolumeService) GetByWorkID(workID uint) ([]model.Volume, error) {
	return s.repos.Volume.FindByWorkID(workID)
}

func (s *VolumeService) Update(id uint, req *request.UpdateVolumeRequest) (*model.Volume, error) {
	volume, err := s.repos.Volume.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		volume.Title = req.Title
	}
	if req.Summary != "" {
		volume.Summary = req.Summary
	}

	if err := s.repos.Volume.Update(volume); err != nil {
		return nil, err
	}

	return volume, nil
}

func (s *VolumeService) Delete(id uint) error {
	return s.repos.Volume.Delete(id)
}

type ChapterService struct {
	repos *repository.Repositories
}

func NewChapterService(repos *repository.Repositories) *ChapterService {
	return &ChapterService{repos: repos}
}

func (s *ChapterService) Create(workID uint, req *request.CreateChapterRequest) (*model.Chapter, error) {
	chapter := &model.Chapter{
		WorkID:       workID,
		VolumeID:     req.VolumeID,
		ChapterIndex: req.ChapterIndex,
		Title:        req.Title,
		Summary:      req.Summary,
		TimeSpace:    req.TimeSpace,
		Characters:   req.Characters,
		Scenes:       req.Scenes,
	}

	if err := s.repos.Chapter.Create(chapter); err != nil {
		return nil, err
	}

	return chapter, nil
}

func (s *ChapterService) GetByWorkID(workID uint) ([]model.Chapter, error) {
	return s.repos.Chapter.FindByWorkID(workID)
}

func (s *ChapterService) GetByID(id uint) (*model.Chapter, error) {
	return s.repos.Chapter.FindByID(id)
}

func (s *ChapterService) Update(id uint, req *request.UpdateChapterRequest) (*model.Chapter, error) {
	chapter, err := s.repos.Chapter.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		chapter.Title = req.Title
	}
	if req.Summary != "" {
		chapter.Summary = req.Summary
	}

	if err := s.repos.Chapter.Update(chapter); err != nil {
		return nil, err
	}

	return chapter, nil
}

func (s *ChapterService) Delete(id uint) error {
	return s.repos.Chapter.Delete(id)
}
