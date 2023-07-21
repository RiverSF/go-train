package crawl

import (
	"fmt"
	colly "github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"river/gomod/excel"
	"strings"
	"time"
)

func GoColly() {
	//colly的主体是Collector对象，管理网络通信和负责在作业运行时执行附加的回掉函数
	c := colly.NewCollector(
		// 开启本机debug
		colly.Debugger(&debug.LogDebugger{}),
	)

	//发送请求之前的执行函数
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("这里是发送之前执行的函数")
	})

	//发送请求错误被回调
	c.OnError(func(r *colly.Response, err error) {
		fmt.Print(err)
	})

	//响应请求之后被回调
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
		fmt.Println("Response body length：", len(r.Body))
		r.Save("./crawl/movie_douban.html")
	})

	//response之后会调用该函数 OnHTML，分析页面数据
	/*正在热映*/
	var hotMovies [][]string
	hotMovies = append(hotMovies, []string{"电影名字", "豆瓣电影详情链接", "电影封面", "评分"})
	c.OnHTML(".ui-slide-item", func(e *colly.HTMLElement) {
		e.ForEach("ul", func(_ int, el *colly.HTMLElement) {
			hot_movie_name := e.ChildAttr(".poster a img", "alt")
			hot_movie_url := e.ChildAttr(".poster a", "href")
			hot_movie_img := e.ChildAttr(".poster a img", "src")
			score := e.ChildText(".rating .subject-rate")
			//fmt.Println(hot_movie_url, hot_movie_img, hot_movie_name, score)

			hotMovies = append(hotMovies, []string{hot_movie_name, hot_movie_url, hot_movie_img, score})
		})
	})

	/*最近热门电影*/
	c.OnHTML(".gaia-movie .tag-list", func(e *colly.HTMLElement) {
	})

	/*导航*/
	nav := make(map[string]string)
	c.OnHTML(".nav-items ul", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			navTitle := el.ChildText("a")
			navUrl := el.ChildAttr("a", "href")
			nav[navTitle] = navUrl
		})
	})

	/*热门影评*/
	type reviewS struct {
		title, person, personPage, content, movie, movieUrl string
	}
	reviews := make(map[string]reviewS)
	c.OnHTML("#reviews", func(e *colly.HTMLElement) {
		e.ForEach(".review ", func(_ int, el *colly.HTMLElement) {

			title := el.ChildText(".review-bd h3 a")
			content := el.ChildText(".review-bd div.review-content")
			content = strings.Fields(content)[0]

			reviewMeta := make(chan string, 4)
			el.ForEach(".review-bd div.review-meta a", func(_ int, ele *colly.HTMLElement) {
				reviewMeta <- ele.Attr("href")
				reviewMeta <- ele.Text
			})

			personPage, person, movieUrl, movie := <-reviewMeta, <-reviewMeta, <-reviewMeta, <-reviewMeta
			r := reviewS{title, person, personPage, content, movie, movieUrl}
			reviews[movie] = r
		})
	})

	//在OnHTML之后被调用
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		fmt.Println(hotMovies)
		fmt.Println(nav)
		fmt.Println(reviews)

		//导出热门电影
		dateTime := time.Now().Format("20060102")
		excel.UpdateExcel(hotMovies, "./download/hot_movies.xlsx", dateTime+"-热门电影")
	})

	//这里是执行访问url
	url := "https://movie.douban.com/"
	c.Visit(url)
}
