package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"personal-assistant/internal/config"
	"personal-assistant/internal/middlewares"
	"personal-assistant/internal/repositories"
	"personal-assistant/internal/wire"
)

func Router(cfg config.Config) *gin.Engine {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "x-api-key"}
	router.Use(cors.New(corsConfig))

	db := repositories.NewDatabase()
	database := db.Connect()

	// controller
	whisperController := wire.InitializeWhisperController(cfg, database)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	apiRoute := router.Group("/api")
	apiRoute.Use(middlewares.APIKeyMiddleware())
	{
		whisperRoute := apiRoute.Group("/whisper")
		{
			whisperController.RegisterRoutes(whisperRoute)
		}
	}

	return router
}
