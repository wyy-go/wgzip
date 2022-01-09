# wgzip

![GitHub Repo stars](https://img.shields.io/github/stars/wyy-go/wgzip?style=social)
![GitHub](https://img.shields.io/github/license/wyy-go/wgzip)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wyy-go/wgzip)
![GitHub CI Status](https://img.shields.io/github/workflow/status/wyy-go/wgzip/ci?label=CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/wyy-go/wgzip)](https://goreportcard.com/report/github.com/wyy-go/wgzip)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/wyy-go/wgzip?tab=doc)
[![codecov](https://codecov.io/gh/wyy-go/wgzip/branch/main/graph/badge.svg)](https://codecov.io/gh/wyy-go/wgzip)

Gin middleware to enable `GZIP` support.

## Usage

Download and install it:

```sh
go get github.com/wyy-go/wgzip
```

Import it in your code:

```go
import "github.com/wyy-go/wgzip"
```

Canonical example:

```go
package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/wyy-go/wgzip"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(wgzip.New(wgzip.WithCompressionType(wgzip.DefaultCompression)))
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

Customized Excluded Extensions

```go
package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/wyy-go/wgzip"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(wgzip.New(wgzip.WithCompressionType(wgzip.DefaultCompression), wgzip.WithExcludedExtensions([]string{".pdf", ".mp4"})))
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

Customized Excluded Paths

```go
package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/wyy-go/wgzip"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(wgzip.New(wgzip.WithCompressionType(wgzip.DefaultCompression), wgzip.WithExcludedPaths([]string{"/api/"})))
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

Customized Excluded Paths

```go
package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/wyy-go/wgzip"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(wgzip.New(wgzip.WithCompressionType(wgzip.DefaultCompression), wgzip.WithExcludedPathsRegexs([]string{".*"})))
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```
