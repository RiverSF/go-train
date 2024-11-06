package net

import (
	"bytes"
	"compress/gzip"
	"sync"
)

// 对象缓存池
type gzipWriterPool struct {
	pool *sync.Pool
}

var GzipWriterPool = &gzipWriterPool{
	pool: &sync.Pool{
		New: func() interface{} {
			buf := new(bytes.Buffer)
			return gzip.NewWriter(buf)
		},
	},
}

func (p *gzipWriterPool) Get() *gzip.Writer {
	w := p.pool.Get().(*gzip.Writer)
	return w
}

func (p *gzipWriterPool) Put(w *gzip.Writer) {
	p.pool.Put(w)
}
