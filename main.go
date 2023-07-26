package main

import (
	"fmt"
	"river/gomod/crawl"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

type Struct struct {
	X, Y int
	Z, w string
}

var arr [10]int

func main() {
	//arr[0] = 1
	//
	//fmt.Println("数组 = ", arr)
	//fmt.Println("数组长度 = ", len(arr))
	//fmt.Println("数组容量 = ", cap(arr))
	//
	//year, month, day := time.Now().Date()
	date_time := time.Now().Format("20060102")
	fmt.Println("Now Time ", date_time)
	//
	//fmt.Println("My favorite number is", rand.Int31n(1000))
	//
	//aaa := "中国"
	//fmt.Println(len(aaa), utf8.RuneCountInString(aaa), []rune(aaa))
	//
	////string转换为int
	//bbb := "123"
	//intB, _ := strconv.Atoi(bbb)
	//fmt.Println(intB)
	//
	//float64 转换成 string
	//strconv.FormatFloat(starF, 'f', 1, 32)
	//
	////defer 在跳出当前方法时执行，先进后出执行
	//fmt.Println(c())
	//
	////指针 空值
	//var ptr *int
	//if ptr == nil {
	//	fmt.Printf("ptr 的值为 : %p\n", ptr)
	//}
	//
	////wc.Test(WordCount)
	//
	////ArrayToString()
	//bytes := [4]byte{1,2,3,4}
	//str := SliceTypeChange(bytes[:])
	//fmt.Println(str)

	// 爬虫训练
	crawl.CrawlTrain()
	//crawl.CrawlDemo()
	//crawl.GoColly()

	// Excel处理训练
	//excel.CreateExcel()
	//excel.ReadExcel()
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

// ArrayToString 定长数组转换为字符串
func ArrayToString() {
	s := [4]byte{65, 66, 67, 68}
	//第一步：先将定长数组转换成切片
	ss := s[:]
	//方式一：标准转换
	aa := string(ss)
	fmt.Println(aa)

	//方式二：强转换 黑魔法
	//https://zhuanlan.zhihu.com/p/270626496
	bb := *(*string)(unsafe.Pointer(&ss))
	fmt.Println(bb)
}

// SliceTypeChange 切片类型之间的转换
func SliceTypeChange(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ".")
}

func WordCount(s string) map[string]int {

	//字符串 转换为 切片
	//var ss []string = strings.Fields(s)
	ss := strings.Fields(s)

	//var m map[string]int
	m := make(map[string]int)

	for _, v := range ss {
		count, ok := m[v]
		if ok {
			m[v] = count + 1
		} else {
			m[v] = 1
		}
	}
	return m
}
