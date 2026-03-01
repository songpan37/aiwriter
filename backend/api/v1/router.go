package v1

import (
	"aiwriter/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, h *handler.Handler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}

	works := r.Group("/works")
	{
		works.GET("", h.GetWorks)
		works.POST("", h.CreateWork)
		works.GET("/:id", h.GetWork)
		works.PUT("/:id", h.UpdateWork)
		works.DELETE("/:id", h.DeleteWork)
		works.GET("/:workId/volumes", h.GetVolumes)
		works.POST("/:workId/volumes", h.CreateVolume)
		works.GET("/:workId/chapters", h.GetChapters)
		works.POST("/:workId/chapters", h.CreateChapter)
	}

	r.GET("/profile", h.GetProfile)
}
