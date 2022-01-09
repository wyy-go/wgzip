package wgzip

import (
	"compress/gzip"
	"io/ioutil"
	"sync"
)

type gzPool struct {
	pool            sync.Pool
	compressionType CompressionType
}

type Pool interface {
	Get() *gzip.Writer
	Put(w *gzip.Writer)
}

func newGzPool(compressionType CompressionType) Pool {
	return &gzPool{compressionType: compressionType}
}

func (p *gzPool) Get() *gzip.Writer {
	w := p.pool.Get()
	if w == nil {
		w, err := gzip.NewWriterLevel(ioutil.Discard, int(p.compressionType))
		if err != nil {
			panic(err)
		}
		return w
	}
	return w.(*gzip.Writer)
}

func (p *gzPool) Put(w *gzip.Writer) {
	p.pool.Put(w)
}
