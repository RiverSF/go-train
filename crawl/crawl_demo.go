package crawl

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sync"
	"time"
	"unsafe"
)

type urlContent struct {
	title, content string
	otherUrl       []string
}

// urlCounter1 记录已抓取的URL
type urlCounter1 struct {
	url map[string]int
	mux sync.Mutex
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func handle(url string, depth int, counter *urlCounter1) {
	fmt.Println(url, depth)

	if depth <= 0 {
		return
	}

	// 已请求过的 URL 去重
	_, ok := counter.url[url]
	fmt.Println(ok)
	if ok {
		return
	}
	counter.mux.Lock()
	counter.url[url] = 1
	counter.mux.Unlock()
	fmt.Println(counter.url)

	// 请求 URL
	err := crawl(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	//并发抓取
	//fmt.Printf("found: %s %q %v\n", urls)
	//for _, u := range urls {
	//	go handle(u, depth-1, counter)
	//}

	time.Sleep(1 * time.Millisecond)
	fmt.Println("end")
	return
}

func crawl(url string) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("【%s】 Fatal Error: %s", url, err.Error())
	}

	//设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("【%s】 Request Error: %s", url, err.Error())
	}
	fmt.Println(resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("【%s】 Body Error: %s", url, err.Error())
	}

	//字节流转换为字符串
	html := *(*string)(unsafe.Pointer(&body))

	fmt.Println(html)

	reg := regexp.MustCompile("<title>.+</title>")
	title := reg.FindAllStringSubmatch(html, -1)

	fmt.Println(title)

	return nil
}

// CrawlDemo 抓取豆瓣
func CrawlDemo() {
	//counter := &urlCounter1{url: make(map[string]int)}
	//handle("https://golang.org/", 4, counter)
	crawl("https://movie.douban.com/")
}
