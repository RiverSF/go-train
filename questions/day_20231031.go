package questions

import "fmt"

func app() func(string) string {
	t := "Hi"
	fmt.Println("&t = ", &t)
	c := func(b string) string {
		t = t + " " + b
		fmt.Println("&t =", &t, "t = ", t)
		return t
	}
	return c
}

func Q20231031() {
	/**
	https://studygolang.com/interview/question

	以下代码输出什么？
	如果最后再加一行代码：fmt.Println(a("All"))，它输出什么？
	（提示：你可以输出 t 的地址，看看是什么情况。）
	*/

	a := app()
	b := app()
	a("A")
	fmt.Println(b("B"), &b)
	fmt.Println(a("All"), &a)
}

func q20231031() {
	// 什么是闭包

	// 在支持函数是一等公民的语言中，一个函数的返回值是另一个函数，被返回的函数可以访问父函数内的变量，当这个被返回的函数在外部执行时，就产生了闭包。

	// 总结：
	//对闭包来说，函数在该语言中得是一等公民。一般来说，一个函数返回另外一个函数，这个被返回的函数可以引用外层函数的局部变量，这形成了一个闭包。
	//通常，闭包通过一个结构体来实现，它存储一个函数和一个关联的上下文环境。
	//但 Go 语言中，匿名函数就是一个闭包，它可以直接引用外部函数的局部变量，因为 Go 规范和 FAQ 都这么说了。
}
