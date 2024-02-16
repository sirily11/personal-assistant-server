package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sme-demo/internal/config"
	"sme-demo/internal/middlewares"
)

func Router(config config.Config) *gin.Engine {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "x-api-key"}
	router.Use(cors.New(corsConfig))

	// controller

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	_ = router.Group("/api").Use(middlewares.APIKeyMiddleware())
	{

	}
	return router
}
