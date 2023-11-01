package train

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex

	// 结构体可以自动继承匿名内部结构体的所有方法
	sync.Mutex //互斥锁
	//sync.RWMutex//读写锁
}

// 使用结构体中匿名方法
func (c *SafeCounter) Sub(key string) {
	c.Lock()
	c.v[key]--
	c.Unlock()
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v

	// 在 return 之后解锁
	defer c.mux.Unlock()

	return c.v[key]
}

func MutexOut() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
