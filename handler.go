package wgzip

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

type gzipHandler struct {
	*Options
}

func New(options ...Option) gin.HandlerFunc {
	return newGzipHandler(options...).Handle
}

func newGzipHandler(opts ...Option) *gzipHandler {
	options := newOptions(opts...)

	handler := &gzipHandler{
		Options: options,
	}

	return handler
}

func (h *gzipHandler) Handle(c *gin.Context) {
	if fn := h.decompressFn; fn != nil &&
		c.Request.Header.Get("Content-Encoding") == "gzip" {
		fn(c)
	}

	if !h.shouldCompress(c.Request) {
		return
	}

	w := h.pool.Get()
	defer h.pool.Put(w)
	defer w.Reset(ioutil.Discard)
	w.Reset(c.Writer)

	c.Header("Content-Encoding", "gzip")
	c.Header("Vary", "Accept-Encoding")
	c.Writer = &gzipWriter{c.Writer, w}
	defer func() {
		w.Close()
		c.Header("Content-Length", fmt.Sprint(c.Writer.Size()))
	}()
	c.Next()
}

func (h *gzipHandler) shouldCompress(req *http.Request) bool {
	if !strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") ||
		strings.Contains(req.Header.Get("Connection"), "Upgrade") ||
		strings.Contains(req.Header.Get("Accept"), "text/event-stream") {

		return false
	}

	extension := filepath.Ext(req.URL.Path)

	if h.excludedExtensions.Contains(extension) {
		return false
	}

	if h.excludedPaths.Contains(req.URL.Path) {
		return false
	}
	if h.excludedPathesRegexs.Contains(req.URL.Path) {
		return false
	}

	return true
}
