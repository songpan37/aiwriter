package handler

import (
	"net/http"
	"strconv"

	"aiwriter/internal/dto/request"
	"aiwriter/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService    *service.AuthService
	workService    *service.WorkService
	volumeService  *service.VolumeService
	chapterService *service.ChapterService
	repos          *service.Service
}

func NewHandlers(services *service.Service) *Handler {
	return &Handler{
		authService:    service.NewAuthService(services.Repos, services.Config),
		workService:    service.NewWorkService(services.Repos, services.Store),
		volumeService:  service.NewVolumeService(services.Repos),
		chapterService: service.NewChapterService(services.Repos, services.Store),
		repos:          services,
	}
}

func (h *Handler) GetUserID(c *gin.Context) uint {
	userID, _ := c.Get("userID")
	return userID.(uint)
}

func (h *Handler) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1002,
			"message": err.Error(),
		})
		return
	}

	result, err := h.authService.Register(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    result,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1002,
			"message": err.Error(),
		})
		return
	}

	result, err := h.authService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    2001,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    result,
	})
}

func (h *Handler) GetProfile(c *gin.Context) {
	userID := h.GetUserID(c)
	user, err := h.repos.Repos.User.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"avatar":     user.AvatarKey,
			"created_at": user.CreatedAt,
		},
	})
}

func (h *Handler) CreateWork(c *gin.Context) {
	var req request.CreateWorkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1002,
			"message": err.Error(),
		})
		return
	}

	userID := h.GetUserID(c)
	work, err := h.workService.Create(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    0,
		"message": "success",
		"data":    work,
	})
}

func (h *Handler) GetWorks(c *gin.Context) {
	userID := h.GetUserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	result, err := h.workService.GetList(userID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    result,
	})
}

func (h *Handler) GetWork(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	work, err := h.workService.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    3001,
			"message": "Work not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    work,
	})
}

func (h *Handler) UpdateWork(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	var req request.UpdateWorkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1002,
			"message": err.Error(),
		})
		return
	}

	work, err := h.workService.Update(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    work,
	})
}

func (h *Handler) DeleteWork(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := h.workService.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"code":    0,
		"message": "success",
	})
}

func (h *Handler) CreateVolume(c *gin.Context) {
	workID, _ := strconv.ParseUint(c.Param("workId"), 10, 32)

	var req request.CreateVolumeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1002,
			"message": err.Error(),
		})
		return
	}

	volume, err := h.volumeService.Create(uint(workID), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    0,
		"message": "success",
		"data":    volume,
	})
}

func (h *Handler) GetVolumes(c *gin.Context) {
	workID, _ := strconv.ParseUint(c.Param("workId"), 10, 32)

	volumes, err := h.volumeService.GetByWorkID(uint(workID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    volumes,
	})
}

func (h *Handler) CreateChapter(c *gin.Context) {
	workID, _ := strconv.ParseUint(c.Param("workId"), 10, 32)

	var req request.CreateChapterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1002,
			"message": err.Error(),
		})
		return
	}

	chapter, err := h.chapterService.Create(uint(workID), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    0,
		"message": "success",
		"data":    chapter,
	})
}

func (h *Handler) GetChapters(c *gin.Context) {
	workID, _ := strconv.ParseUint(c.Param("workId"), 10, 32)

	chapters, err := h.chapterService.GetByWorkID(uint(workID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1000,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    chapters,
	})
}
