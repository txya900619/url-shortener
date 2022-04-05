package httperr

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/txya900619/url-shortener/pkg/errors"
)

func InternalError(slug string, err error, c *gin.Context) {
	httpRespondWithError(err, slug, c, "Internal server error", http.StatusInternalServerError)
}

func NotFound(slug string, err error, c *gin.Context) {
	httpRespondWithError(err, slug, c, "Not found", http.StatusNotFound)
}

func BadRequest(slug string, err error, c *gin.Context) {
	httpRespondWithError(err, slug, c, "Bad request", http.StatusBadRequest)
}

func RespondWithSlugError(err error, c *gin.Context) {
	slugError, ok := err.(errors.SlugError)
	if !ok {
		InternalError("internal-server-error", err, c)
		return
	}

	switch slugError.ErrorType() {
	case errors.ErrorTypeNotFound:
		NotFound(slugError.Slug(), slugError, c)
	case errors.ErrorTypeIncorrectInput:
		BadRequest(slugError.Slug(), slugError, c)
	default:
		InternalError(slugError.Slug(), slugError, c)
	}
}

func httpRespondWithError(err error, slug string, c *gin.Context, logMSg string, status int) {
	logrus.WithError(err).WithField("error-slug", slug).Warn(logMSg)

	c.AbortWithStatusJSON(status, &ErrorResponse{Slug: slug})
}

type ErrorResponse struct {
	Slug string `json:"slug"`
}
