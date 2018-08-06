package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// DefaultHeaderName default header name
	DefaultHeaderName = "X-Health-Check"

	// DefaultHeaderValue default header value
	DefaultHeaderValue = "1"

	// DefaultResponseCode default response code
	DefaultResponseCode = http.StatusOK

	// DefaultResponseText default response text
	DefaultResponseText = "ok"

	// DefaultConfig default config
	DefaultConfig = Config{
		HeaderName:   DefaultHeaderName,
		HeaderValue:  DefaultHeaderValue,
		ResponseCode: DefaultResponseCode,
		ResponseText: DefaultResponseText}
)

// Config holds the configuration values
type Config struct {
	HeaderName   string
	HeaderValue  string
	ResponseCode int
	ResponseText string
}

// Default creates a new middileware with the default configuration
func Default() gin.HandlerFunc {
	return New(DefaultConfig)
}

// New creates a new middileware with the `cfg`
func New(cfg Config) gin.HandlerFunc {
	if cfg.HeaderName == "" {
		cfg.HeaderName = DefaultHeaderName
	}
	if cfg.HeaderValue == "" {
		cfg.HeaderValue = DefaultHeaderValue
	}
	if cfg.ResponseCode == 0 {
		cfg.ResponseCode = DefaultResponseCode
	}

	return func(ctx *gin.Context) {
		if ctx.GetHeader(cfg.HeaderName) == cfg.HeaderValue {
			ctx.String(cfg.ResponseCode, cfg.ResponseText)
			ctx.Abort()
		}

		ctx.Next()
	}
}
