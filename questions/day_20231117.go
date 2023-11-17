package questions

import "fmt"

const (
	a = iota
	b = iota
)

const (
	c = 'c'
	d = iota
	e = 50
	ff
)

func Q20231117() {
	fmt.Println(a, b, c, d, e, ff)
}

func q20231117() {
	/**
	知识点：iota

	iota 是 golang 语言的常量计数器，只能在常量表达式中使用

	iota 在 const 关键字出现时，将被重置为0， const 中每新增一行常量声明将使iota计数一次

	详细解析见： https://studygolang.com/articles/2192
	*/
}
