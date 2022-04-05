package ports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/txya900619/url-shortener/internal/shorturl/app"
	"github.com/txya900619/url-shortener/pkg/ginx/httperr"
)

type HttpServer struct {
	service app.ShortUrlService
}

func NewHttpServer(service app.ShortUrlService) *HttpServer {
	return &HttpServer{
		service: service,
	}
}

func (h HttpServer) CreateShortUrl(c *gin.Context) {
	var postShortUrl PostShortUrl

	if err := c.BindJSON(&postShortUrl); err != nil {
		httperr.BadRequest("invalid-request", err, c)
		return
	}

	shortUrlId, err := h.service.CreateShortUrl(c, postShortUrl.ExpireAt, postShortUrl.Url)
	if err != nil {
		httperr.RespondWithSlugError(err, c)
		return
	}

	c.JSON(http.StatusCreated, shortUrlIdToResponse(shortUrlId))
}

func shortUrlIdToResponse(shortUrlId string) ShortUrl {
	return ShortUrl{
		Id:       shortUrlId,
		ShortUrl: "http://localhost:8080/" + shortUrlId, //TODO: host add to env
	}
}

func (h HttpServer) RedirectToOriginUrl(c *gin.Context, urlId string) {
	longUrl, err := h.service.RedirectToOriginUrl(c, urlId)
	if err != nil {
		httperr.RespondWithSlugError(err, c)
		return
	}

	c.Redirect(http.StatusFound, longUrl)
}
