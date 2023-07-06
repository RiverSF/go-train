package train

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// CrawlUrlCounter 记录已抓取的URL
type CrawlUrlCounter struct {
	url map[string]int
	mux sync.Mutex
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, counter *CrawlUrlCounter) {
	fmt.Println(url, depth)

	if depth <= 0 {
		return
	}

	_, ok := counter.url[url]
	fmt.Println(ok)
	if ok {
		return
	}
	counter.mux.Lock()
	counter.url[url] = 1
	counter.mux.Unlock()
	fmt.Println(counter.url)

	body, urls, err := fetcher.Fetch(url)
	//_, urls, err := fetcher.Fetch(url)
	if err != nil {
		//fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q %v\n", url, body, urls)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, counter)
		//time.Sleep(1 * time.Millisecond)
	}

	fmt.Println("end")
	//return
}

func CrawlPage() {
	counter := &CrawlUrlCounter{url: make(map[string]int)}
	Crawl("https://golang.org/", 4, fetcher, counter)
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
