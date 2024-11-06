package main

import (
	"fmt"
	"os"
	"river/train"
	"strconv"
	"strings"
	"sync"
	"unsafe"
)

type Struct struct {
	X, Y int
	Z, w string
}

type Test struct {
	Ac string `json:"aa"`           //字段在JSON中显示为键 a
	Bc string `json:"bb,omitempty"` //字段在JSON中显示为键 b, 如果该字段的值为空，则从对象中省略该字段
	Cc string `json:",omitempty"`   //字段在JSON中显示为键"C"(默认值)，但是如果字段值为空，则跳过该字段
	Dc string `json:"-"`            //忽略该字段
	Ec int    `json:"-,"`           //字段在JSON中显示为键"-"
}

var mu sync.Mutex

//基本数据类型 区分
//值类型： int 系列, float 系列, bool, string 、数组和结构体 struct
//引用类型：指针、slice 切片、map、管道 chan、interface、函数 都是引用类型

func init() {
	/**
	init() 函数是 Go 程序初始化的一部分。Go 程序初始化先于 main 函数，由 runtime 初始化每个导入的包，初始化顺序不是按照从上到下的导入顺序，而是按照解析的依赖关系，没有依赖的包最先初始化。

	每个包首先初始化包作用域的常量和变量（常量优先于变量），然后执行包的 init() 函数。同一个包，甚至是同一个源文件可以有多个 init() 函数。init() 函数没有入参和返回值，不能被其他函数调用，同一个包内多个 init() 函数的执行顺序不作保证。

	一句话总结： import –> const –> var –> init() –> main()
	*/
	fmt.Println("init1:", 123)

	//list := []Test{
	//	{
	//		Ec: 2,
	//	},
	//	{
	//		Ec: 1,
	//	},
	//}
	//
	//sort.Slice(list, func(i, j int) bool { // Sort优先级排序
	//	if list[i].Ec < list[j].Ec {
	//		return true
	//	}
	//	return false
	//})

	train.TimeRate()
}

// 完成对变量的拷贝赋值
func (d *Test) clone() *Test {
	template := *d   // * 读取指针值
	return &template // & 生成指针
}

func main() {
	//go 训练题
	//questions.Q20240528()

	//mu.Lock()
	// new 产生一个指针类型变量，建议使用make初始化变量
	//arr := new([]int)
	//mu.Unlock()
	//fmt.Println("数组 = ", *arr)
	//fmt.Println("数组长度 = ", len(arr))
	//fmt.Println("数组容量 = ", cap(arr))

	//stringSpaceFilter()

	//train.ChannelOut()

	//year, month, day := time.Now().Date()
	//date_time := time.Now().Format("20060102")
	//fmt.Println("Now Time ", date_time)

	// json 转义
	//t := Test{
	//	"abcde",
	//	"abcde",
	//	"abcde",
	//	"abcde",
	//	"abcde",
	//}
	//en_t, _ := json.Marshal(t)	//返回值为字节切片，string 类型转换后可打印明文内容
	//en_t_str := string(en_t)
	//fmt.Println(en_t_str)
	// json 解析
	//var m Test
	//json.Unmarshal(en_t, &m)
	//json.Unmarshal([]byte(en_t_str), &m)
	//fmt.Println(m, m.Dc12)

	//fmt.Println("My favorite number is", rand.Int31n(1000))
	//
	//aaa := "中国"
	//fmt.Println(len(aaa), utf8.RuneCountInString(aaa), []rune(aaa))
	//
	//string转换为int
	//bbb := "123"
	//intB, _ := strconv.Atoi(bbb)
	//fmt.Println(intB)
	//
	//float64 转换成 string
	//strconv.FormatFloat(starF, 'f', 1, 32)
	//
	//defer 在跳出当前方法时执行，先进后出执行
	//fmt.Println(c())
	//
	//指针 空值
	//var ptr *int
	//if ptr == nil {
	//	fmt.Printf("ptr 的值为 : %p\n", ptr)
	//}
	//
	//wc.Test(WordCount)
	//
	//ArrayToString()
	//bytes := [4]byte{1,2,3,4}
	//str := SliceTypeChange(bytes[:])
	//fmt.Println(str)

	// 爬虫训练
	//crawl.CrawlTrain()
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

// html 实体空格回车过滤
func stringSpaceFilter() {

	txt, _ := os.ReadFile("strings.txt")
	fmt.Println(txt)

	oStr := string(txt)
	fmt.Println(oStr)

	utf8Spaces := []rune{194, 160}
	for _, v := range utf8Spaces {
		oStr = strings.ReplaceAll(oStr, string(v), "")
	}
	oStr = strings.ReplaceAll(oStr, " ", "")
	fmt.Println([]byte(oStr))
	fmt.Println(oStr)
}
