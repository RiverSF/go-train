package net

import (
	"bytes"
	"sync"
)

type bufferPool struct {
	pool *sync.Pool
}

var BufferPool = &bufferPool{
	pool: &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	},
}

func (p *bufferPool) Get() *bytes.Buffer {
	w := p.pool.Get().(*bytes.Buffer)
	return w
}

func (p *bufferPool) Put(w *bytes.Buffer) {
	w.Reset()
	p.pool.Put(w)
}
