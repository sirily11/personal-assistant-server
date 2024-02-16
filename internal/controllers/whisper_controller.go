package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "sme-demo/internal/repositories/model/whisper"
	. "sme-demo/internal/services"
)

type WhisperController struct {
	whisperService WhisperModelServiceInterface
}

func NewWhisperController(service WhisperModelServiceInterface) *WhisperController {
	return &WhisperController{
		whisperService: service,
	}
}

func (controller *WhisperController) Create(c *gin.Context) {
	whisper := WhisperModel{}
	Validate[WhisperModel](c, &whisper)

	created, err := controller.whisperService.Create(whisper)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (controller *WhisperController) List(c *gin.Context) {
	list, err := controller.whisperService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list)
}

func (controller *WhisperController) GetById(c *gin.Context) {
	id := c.Param("id")
	get, err := controller.whisperService.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, get)
}

func (controller *WhisperController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := controller.whisperService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (controller *WhisperController) Update(c *gin.Context) {
	id := c.Param("id")
	whisper := WhisperModel{}
	Validate[WhisperModel](c, &whisper)

	updated, err := controller.whisperService.Update(id, whisper)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (controller *WhisperController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("", controller.Create)
	router.GET("", controller.List)
	router.GET("/:id", controller.GetById)
	router.DELETE("/:id", controller.Delete)
	router.PATCH("/:id", controller.Update)
}
