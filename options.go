package wgzip

import (
	"compress/gzip"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	defaultExcludedExtentions = newExcludedExtensions([]string{
		".png", ".gif", ".jpeg", ".jpg",
	})

	defaultOptions = &Options{
		excludedExtensions: defaultExcludedExtentions,
		compressionType:    DefaultCompression,
		decompressFn:       defaultDecompressHandle,
	}
)

type CompressionType int
type DecompressFn func(c *gin.Context)

const (
	BestCompression    CompressionType = gzip.BestCompression
	BestSpeed          CompressionType = gzip.BestSpeed
	DefaultCompression CompressionType = gzip.DefaultCompression
	NoCompression      CompressionType = gzip.NoCompression
)

type Options struct {
	excludedExtensions   excludedExtensions
	excludedPaths        excludedPaths
	excludedPathesRegexs excludedPathesRegexs
	decompressFn         DecompressFn
	compressionType      CompressionType
	pool                 Pool
}

type Option func(*Options)

func WithExcludedExtensions(args []string) Option {
	return func(o *Options) {
		o.excludedExtensions = newExcludedExtensions(args)
	}
}

func WithExcludedPaths(args []string) Option {
	return func(o *Options) {
		o.excludedPaths = newExcludedPaths(args)
	}
}

func WithExcludedPathsRegexs(args []string) Option {
	return func(o *Options) {
		o.excludedPathesRegexs = newExcludedPathesRegexs(args)
	}
}

func WithDecompressFn(decompressFn DecompressFn) Option {
	return func(o *Options) {
		o.decompressFn = decompressFn
	}
}

func WithCompressionType(compressionType CompressionType) Option {
	return func(o *Options) {
		o.compressionType = compressionType
	}
}

func WithPool(pool Pool) Option {
	return func(o *Options) {
		o.pool = pool
	}
}

// Using map for better lookup performance
type excludedExtensions map[string]bool

func newExcludedExtensions(extensions []string) excludedExtensions {
	res := make(excludedExtensions)
	for _, e := range extensions {
		res[e] = true
	}
	return res
}

func (e excludedExtensions) Contains(target string) bool {
	_, ok := e[target]
	return ok
}

type excludedPaths []string

func newExcludedPaths(paths []string) excludedPaths {
	return excludedPaths(paths)
}

func (e excludedPaths) Contains(requestURI string) bool {
	for _, path := range e {
		if strings.HasPrefix(requestURI, path) {
			return true
		}
	}
	return false
}

type excludedPathesRegexs []*regexp.Regexp

func newExcludedPathesRegexs(regexs []string) excludedPathesRegexs {
	result := make([]*regexp.Regexp, len(regexs))
	for i, reg := range regexs {
		result[i] = regexp.MustCompile(reg)
	}
	return result
}

func (e excludedPathesRegexs) Contains(requestURI string) bool {
	for _, reg := range e {
		if reg.MatchString(requestURI) {
			return true
		}
	}
	return false
}

func defaultDecompressHandle(c *gin.Context) {
	if c.Request.Body == nil {
		return
	}
	r, err := gzip.NewReader(c.Request.Body)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Request.Header.Del("Content-Encoding")
	c.Request.Header.Del("Content-Length")
	c.Request.Body = r
}

func newOptions(opts ...Option) *Options {
	options := defaultOptions
	options.pool = newGzPool(options.compressionType)
	for _, setter := range opts {
		setter(options)
	}
	return options
}
