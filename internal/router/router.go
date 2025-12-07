package router

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"talkerinos/internal/handler"
)

func New(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:3000"
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(allowedOrigins, ","),
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	v1 := router.Group("/v1")
	{
		v1.GET("/blog", h.ListPosts)
		v1.GET("/blog/drafts", h.ListDrafts)
		v1.GET("/blog/:id", h.GetPost)
		v1.GET("/blog/slug/:slug", h.GetPostBySlug)
		v1.POST("/blog", h.AddPosts)
		v1.PUT("/blog/:id", h.UpdatePost)
		v1.DELETE("/blog/:id", h.DeletePost)
	}

	return router
}
