package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func NewEngine() *gin.Engine {
	if viper.GetBool("DEBUG") {
		gin.ForceConsoleColor()
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	engine.Use(
		otelgin.Middleware("gin"),
		gin.Logger(),
		gin.Recovery())

	engine.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })
	engine.GET("/metrics", func(c *gin.Context) { promhttp.Handler().ServeHTTP(c.Writer, c.Request) })

	return engine
}
