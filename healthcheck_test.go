package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	res    *httptest.ResponseRecorder
	router *gin.Engine
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setup(t *testing.T, fns ...func()) func(...func()) {
	res = httptest.NewRecorder()

	router = gin.New()
	for _, fn := range fns {
		fn()
	}
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusNoContent, "")
	})

	return func(fns ...func()) {
		for _, fn := range fns {
			fn()
		}
	}
}

func TestNoConfig(t *testing.T) {
	defer setup(t)()

	makeRequestAndAssert(t,
		http.Header{DefaultHeaderName: []string{DefaultHeaderValue}},
		http.StatusNoContent,
		"",
	)
}

func TestNoHeader(t *testing.T) {
	defer setup(t)()

	makeRequestAndAssert(t,
		nil,
		http.StatusNoContent,
		"",
	)
}

func TestDefaultConfig(t *testing.T) {
	defer setup(t, func() {
		router.Use(Default())
	})()

	makeRequestAndAssert(t,
		http.Header{DefaultHeaderName: []string{DefaultHeaderValue}},
		DefaultResponseCode,
		DefaultResponseText,
	)
}

func TestCustomConfig(t *testing.T) {
	defer setup(t, func() {
		router.Use(New(Config{
			HeaderName:   "X-Custom-Header",
			HeaderValue:  "customValue",
			ResponseCode: http.StatusTeapot,
			ResponseText: "teapot",
		}))
	})()

	makeRequestAndAssert(t,
		http.Header{"X-Custom-Header": []string{"customValue"}},
		http.StatusTeapot,
		"teapot",
	)
}

func TestCustomConfig_empty(t *testing.T) {
	defer setup(t, func() {
		router.Use(New(Config{}))
	})()

	makeRequestAndAssert(t,
		http.Header{DefaultHeaderName: []string{DefaultHeaderValue}},
		DefaultResponseCode,
		"",
	)
}

func makeRequestAndAssert(t *testing.T, headers http.Header, status int, body string) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header = headers

	router.ServeHTTP(res, req)

	if res.Code != status {
		t.Errorf("expected %d, got %d", status, res.Code)
	}
	if b := res.Body.String(); b != body {
		t.Errorf("expected %q, got %q", body, b)
	}
}
