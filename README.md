# gin-health-check

[![Build Status](https://travis-ci.org/RaMin0/gin-health-check.svg?branch=master)](https://travis-ci.org/RaMin0/gin-health-check)
[![CodeCov](https://codecov.io/gh/RaMin0/gin-health-check/branch/master/graph/badge.svg)](https://codecov.io/gh/RaMin0/gin-health-check)
[![GoDoc](https://godoc.org/github.com/RaMin0/gin-health-check?status.svg)](https://godoc.org/github.com/RaMin0/gin-health-check)
[![License](https://img.shields.io/github/license/RaMin0/gin-health-check.svg)](LICENSE.md)

A health check middleware for [Gin](http://gin-gonic.github.io/gin/).

## Installation

``` bash
$ go get -u github.com/RaMin0/gin-health-check
```

## Usage

### Default Config

``` go
import (
	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(healthcheck.Default())
}
```

```bash
$ curl -iL -XGET -H "X-Health-Check: 1" http://localhost
  # HTTP/1.1 200 OK
  # Content-Length: 2
  # Content-Type: text/plain; charset=utf-8
  #
  # ok
```

### Custom Config

``` go
import (
	"net/http"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(healthcheck.New(healthcheck.Config{
		HeaderName:   "X-Custom-Header",
		HeaderValue:  "customValue",
		ResponseCode: http.StatusTeapot,
		ResponseText: "teapot",
	}))
}
```

```bash
$ curl -iL -XGET -H "X-Custom-Header: customValue" http://localhost
  # HTTP/1.1 418 I'm a teapot
  # Content-Length: 6
  # Content-Type: text/plain; charset=utf-8
  #
  # teapot
```
