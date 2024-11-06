package net

import (
	"compress/gzip"
	"sync"
)

// 对象缓存池
type gzipReaderPool struct {
	pool *sync.Pool
}

var GzipReaderPool = &gzipReaderPool{
	pool: &sync.Pool{
		New: func() interface{} {
			return new(gzip.Reader)
		},
	},
}

func (p *gzipReaderPool) Get() *gzip.Reader {
	w := p.pool.Get().(*gzip.Reader)
	return w
}

func (p *gzipReaderPool) Put(w *gzip.Reader) {
	p.pool.Put(w)
}
