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
}

func RegisterProtectedRoutes(r *gin.RouterGroup, h *handler.Handler) {
	r.GET("/profile", h.GetProfile)

	works := r.Group("/works")
	{
		works.GET("", h.GetWorks)
		works.POST("", h.CreateWork)
		works.GET("/:id", h.GetWork)
		works.PUT("/:id", h.UpdateWork)
		works.DELETE("/:id", h.DeleteWork)
	}

	workVolumes := r.Group("/works/:id/volumes")
	{
		workVolumes.GET("", h.GetVolumes)
		workVolumes.POST("", h.CreateVolume)
	}

	workChapters := r.Group("/works/:id/chapters")
	{
		workChapters.GET("", h.GetChapters)
		workChapters.POST("", h.CreateChapter)
	}
}
