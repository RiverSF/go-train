package crawl

import (
	"fmt"
	colly "github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"river/gomod/excel"
	"strconv"
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
		dateTime := time.Now().Format("20060102")
		r.Save("./download/douban/movie_douban_" + dateTime + ".html")
	})

	//response之后会调用该函数 OnHTML，分析页面数据
	/*正在热映*/
	var hotMovies [][]string
	hotMovies = append(hotMovies, []string{"电影名字", "评分", "Star", "评价人数", "上映时间", "电影时长", "地区", "导演", "主演", "电影简介", "预告视频链接", "购票链接", "豆瓣电影详情链接", "电影封面"})
	c.OnHTML(".ui-slide-item", func(e *colly.HTMLElement) {

		data_title := e.Attr("data-title")
		data_release := e.Attr("data-release")
		data_rate := e.Attr("data-rate")
		data_star := e.Attr("data-star")
		data_rater := e.Attr("data-rater")
		data_trailer := e.Attr("data-trailer")
		data_ticket := e.Attr("data-ticket")
		data_duration := e.Attr("data-duration")
		data_region := e.Attr("data-region")
		data_director := e.Attr("data-director")
		data_actors := e.Attr("data-actors")
		data_intro := e.Attr("data-intro")

		var movie_url, movie_img string
		e.ForEach(".poster", func(_ int, el *colly.HTMLElement) {
			movie_url = e.ChildAttr("a", "href")
			movie_img = e.ChildAttr("a img", "src")
		})
		//fmt.Println(data_title, data_rate, data_star, data_rater, data_release, data_duration, data_region, data_director, data_actors, data_intro, data_trailer, data_ticket, movie_url, movie_img)
		hotMovies = append(hotMovies,
			[]string{data_title, data_rate, data_star, data_rater, data_release, data_duration, data_region,
				data_director, data_actors, data_intro, data_trailer, data_ticket, movie_url, movie_img})

		//e.ForEach("ul", func(_ int, el *colly.HTMLElement) {
		//	hot_movie_name := e.ChildAttr(".poster a img", "alt")
		//	hot_movie_url := e.ChildAttr(".poster a", "href")
		//	hot_movie_img := e.ChildAttr(".poster a img", "src")
		//	score := e.ChildText(".rating .subject-rate")
		//	//fmt.Println(hot_movie_url, hot_movie_img, hot_movie_name, score)
		//	hotMovies = append(hotMovies, []string{hot_movie_name, hot_movie_url, hot_movie_img, score})
		//})
	})

	/*最近热门电影*/
	c.OnHTML(".gaia-movie .tag-list", func(e *colly.HTMLElement) {
	})

	/*取出所有电影链接*/
	all_movies := make([]string, 0)
	c.OnHTML("li.title,td.title", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			movie_url := el.Attr("href")
			all_movies = append(all_movies, movie_url)
		})
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
		title, star, person, personPage, content, movie, movieUrl string
	}
	reviewsMap := make(map[string]reviewS)
	c.OnHTML("#reviews", func(e *colly.HTMLElement) {
		e.ForEach(".review ", func(_ int, el *colly.HTMLElement) {

			title := el.ChildText(".review-bd h3 a")
			content := el.ChildText(".review-bd div.review-content")
			content = strings.Fields(content)[0]

			star := el.ChildAttr(".review-bd .review-meta span", "class")
			starS := strings.Split(star, "allstar")[1]
			starI, _ := strconv.Atoi(starS)
			starF := float64(starI) / 10
			starS = strconv.FormatFloat(starF, 'f', 1, 64)

			reviewMeta := make(chan string, 4)
			el.ForEach(".review-bd div.review-meta a", func(_ int, ele *colly.HTMLElement) {
				reviewMeta <- ele.Attr("href")
				reviewMeta <- ele.Text
			})

			personPage, person, movieUrl, movie := <-reviewMeta, <-reviewMeta, <-reviewMeta, <-reviewMeta
			r := reviewS{
				title:      title,
				person:     person,
				personPage: personPage,
				content:    content,
				movie:      movie,
				movieUrl:   movieUrl,
				star:       starS,
			}
			reviewsMap[movie] = r
		})
	})

	//在OnHTML之后被调用
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		//fmt.Println(hotMovies)
		//fmt.Println(nav)
		//fmt.Println(reviewsMap)

		//导出热门电影
		dateTime := time.Now().Format("20060102")
		excel.UpdateExcel(hotMovies, "./download/hot_movies.xlsx", dateTime+"-热门电影")

		//导出最新评论
		reviewsArr := make([][]string, 0)
		reviewsArr = append(reviewsArr, []string{"电影名称", "电影详情链接", "评分", "评论标题", "评论内容", "豆瓣用户", "用户主页"})
		for _, review := range reviewsMap {
			reviewsArr = append(reviewsArr, []string{review.movie, review.movieUrl, review.star, review.title, review.content, review.person, review.personPage})
		}
		excel.UpdateExcel(reviewsArr, "./download/hot_review.xlsx", dateTime+"-热门影评")

		//并发抓取电影详情
		fmt.Println(len(all_movies))
		if all_movies != nil {
			leng := len(all_movies)
			ch := make(chan string, leng)
			for i := 0; i < leng; i++ {
				go concurrencyCrawlMovieDetail(strconv.Itoa(i), ch)
			}
			for i := 0; i < leng; i++ {
				fmt.Println(<-ch)
			}
		}
	})

	//这里是执行访问url
	url := "https://movie.douban.com/"
	c.Visit(url)
}

// 并发抓取电影详情
func concurrencyCrawlMovieDetail(url1 string, ch chan string) {
	url := "https://www.yzktw.com.cn/post/975680.html"
	fmt.Println("并发测试 " + url1)
	c := colly.NewCollector(
	// 开启本机debug
	//colly.Debugger(&debug.LogDebugger{}),
	)
	//响应请求之后被回调
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			fmt.Println(r.Request.URL, " request error, error code = ", r.StatusCode)
		}

		file_name := r.FileName()
		ch <- file_name
		//fmt.Println(file_name)
		//dateTime := time.Now().Format("20060102")
		//r.Save("./download/douban/movie_douban_"+dateTime+".html")
	})

	c.OnHTML("", func(e *colly.HTMLElement) {
	})

	c.OnScraped(func(r *colly.Response) {
	})
	//这里是执行访问url
	c.Visit(url)
}
