package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgtype"
	"github.com/pkg/errors"
	"github.com/swanden/rentateam/internal/domain"
	"github.com/swanden/rentateam/internal/domain/post/repository"
	"github.com/swanden/rentateam/internal/domain/post/usecase"
	"net/http"
	"time"
)

const dateFormat = "2006-01-02 15:04:05"

type postController struct {
	repo     repository.PostRepository
	useCase  *usecase.PostUseCase
	validate *validator.Validate
}

func newPostController(
	ginHandler *gin.RouterGroup,
	repo repository.PostRepository,
	useCase *usecase.PostUseCase,
	validate *validator.Validate) {

	p := postController{
		repo:     repo,
		useCase:  useCase,
		validate: validate,
	}

	h := ginHandler.Group("/post")
	{
		h.POST("/", p.create)
		h.GET("/", p.all)
	}
}

type createRequest struct {
	Title     string   `json:"title" binding:"required" validate:"required" example:"Post title"`
	Body      string   `json:"body" binding:"required" validate:"required" example:"Post body"`
	Tags      []string `json:"tags" binding:"required" validate:"required" example:"FirstTag,SecondTag"`
	CreatedAt string   `json:"created_at" binding:"required" validate:"datetime=2006-01-02 15:04:05" example:"2021-12-01 15:04:05"`
}

type createResponse struct {
	ID int `json:"id"`
}

// @Summary     Create post
// @Description Create post
// @ID          create
// @Tags  	    posts
// @Accept      json
// @Produce     json
// @Param		post body createRequest true "Post fields"
// @Success     201 {object} createResponse
// @Failure     400 {object} responseError
// @Failure     500 {object} responseError
// @Router      /post [post]
func (p *postController) create(c *gin.Context) {
	var request createRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error(), "invalid request body", "http - v1 - postController - create - ShouldBindJSON")
		return
	}

	if err := p.validate.Struct(request); err != nil {
		e := err.(validator.ValidationErrors)
		errorsResponse(c, http.StatusBadRequest, e, "http - v1 - postController - create - Struct")
		return
	}

	requestDate, err := time.Parse(dateFormat, request.CreatedAt)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "wrong format for createdAt", err.Error(), "http - v1 - postController - create - Parse")
		return
	}
	var createdAt pgtype.Timestamp
	err = createdAt.Set(requestDate)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "wrong format for createdAt", err.Error(), "http - v1 - postController - create - Set")
		return
	}
	createDTO := usecase.CreateDTO{
		Title:     request.Title,
		Body:      request.Body,
		Tags:      request.Tags,
		CreatedAt: createdAt,
	}

	id, err := p.useCase.Create(c.Request.Context(), createDTO)

	if errors.Cause(err) == domain.Error {
		errorResponse(c, http.StatusUnprocessableEntity, err.Error(), err.Error(), "http - v1 - postController - create - Cause")
		return
	}
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Internal server error", err.Error(), "http - v1 - postController - create - Cause")
		return
	}

	c.JSON(http.StatusCreated, createResponse{id})
}

type allResponse struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"createdAt"`
}

// @Summary     Show all posts
// @Description Show all posts
// @ID          all
// @Tags  	    posts
// @Accept      json
// @Produce     json
// @Success     201 {object} allResponse
// @Failure     500 {object} responseError
// @Router      /post [get]
func (p *postController) all(c *gin.Context) {
	posts, err := p.repo.All(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "internal server error", err.Error(), "http - v1 - postController - all - All")
		return
	}

	var response allResponse
	for _, post := range posts {
		respPost := Post{
			ID:        post.ID,
			Title:     post.Title,
			Body:      post.Body,
			Tags:      post.Tags,
			CreatedAt: post.CreatedAt.Time.Format(dateFormat),
		}
		response.Posts = append(response.Posts, respPost)
	}

	c.JSON(http.StatusOK, response)
}
