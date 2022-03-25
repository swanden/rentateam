package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"os"
)

type responseError struct {
	Error string `json:"error" example:"message"`
}

type responseErrors struct {
	Errors []string `json:"error"`
}

func errorResponse(c *gin.Context, code int, msg string, logMsg string, source string) {
	fmt.Fprintf(os.Stderr, source+": %v\n", logMsg)
	c.AbortWithStatusJSON(code, responseError{msg})
}

func errorsResponse(c *gin.Context, code int, errors []validator.FieldError, source string) {
	var response responseErrors
	for _, err := range errors {
		response.Errors = append(response.Errors, err.Error())
	}
	fmt.Fprintf(os.Stderr, source+": %v\n", errors)
	c.AbortWithStatusJSON(code, response)
}
