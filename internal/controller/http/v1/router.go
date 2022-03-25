package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/swanden/rentateam/internal/domain/post/repository"
	"github.com/swanden/rentateam/internal/domain/post/usecase"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swanden/rentateam/docs"
)

// Swagger spec:
// @title       Rentateam blog
// @description Rentateam blog test task
// @version     1.0
// @host        localhost:8000
// @BasePath    /v1
func NewRouter(handler *gin.Engine, repo repository.PostRepository, useCase *usecase.PostUseCase, validate *validator.Validate) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := handler.Group("/v1")
	{
		h.GET("/", index)
		newPostController(h, repo, useCase, validate)
	}
}

type apiInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// @Summary     Show API info
// @Description Show API info
// @ID          index
// @Tags  	    api info
// @Accept      json
// @Produce     json
// @Success     200 {object} apiInfo
// @Router      / [get]
func index(c *gin.Context) {
	info := apiInfo{
		Name:    "API",
		Version: "1.0",
	}

	c.JSON(http.StatusOK, info)
}
