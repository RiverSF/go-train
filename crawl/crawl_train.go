package crawl

import (
	"fmt"
	"sync"
)

// WaitGroup 计数器 控制 并发 goroutine 完成
var wg sync.WaitGroup

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// UrlCounter 记录已抓取的URL
// go语言里的map是线程不安全的，可以使用sync.Mutex来保护map
type UrlCounter struct {
	url map[string]int
	mux sync.Mutex
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, counter *UrlCounter) {
	fmt.Println(url, depth)

	if depth <= 0 {
		wg.Done()
		return
	}

	// 已请求过的 URL 去重
	_, ok := counter.url[url]
	fmt.Println(ok)
	if ok {
		wg.Done()
		return
	}
	counter.mux.Lock()
	counter.url[url] = 1
	counter.mux.Unlock()
	fmt.Println(counter.url)

	// 请求 URL
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	//主协程与递归父协程 计数器-1
	//必须延迟执行，否则主协程不会等待子协程执行
	defer wg.Done()

	fmt.Printf("found: %s %q %v\n", url, body, urls)
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, counter)
	}

	//也可以放在此处，for 循环中的子协程不会阻塞，主协程会立即执行到此
	//defer wg.Done()

	//sleep作用：阻塞主协程，等待子协程执行结束
	//多线程中使用睡眠函数不优雅，且子协程执行多久未知，时间不好设置
	//time.Sleep(1 * time.Millisecond)

	//主协程与递归父协程出口
	fmt.Println("end")
	return
}

func CrawlTrain() {
	counter := &UrlCounter{url: make(map[string]int)}
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, counter)
	wg.Wait()
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
