package questions

import (
	"fmt"
)

func Q20231107() {
	/**
	以下代码输出什么？
	*/

	var ans float64 = 15 + 25 + 5.2
	fmt.Println(ans)
}

func q20231107() {

	//常量表达式是指仅包含常量操作数，且是在编译的时候进行计算的
	//而常量，在 Go 语言中又可以分为无类型常量和有类型常量，也可以分为字面值常量和具名常量。
	/**
	const a = 1 + 2 			// a == 3，是无类型常量
	const b int8 = 1 + 2 		// b == 3，是有类型常量，类型是 int8

	而 1、2 这样的就是字面值常量
	a、b 这样的就是具名常量
	*/

	/**
	一个字符串字面量的默认类型是 string 类型。
	一个布尔字面量的默认类型是 bool 类型。
	一个整数型字面量的默认类型是 int 类型。
	一个 rune 字面量的默认类型是 rune（也就是 int32）类型。
	一个浮点数字面量的默认类型是 float64 类型。
	如果一个字面量含有虚部字面量，则此字面量的默认类型是 complex128 类型。
	*/

	//在 Go 语言规范中提到，任何在无类型常量上的操作结果是同一个类别的无类型常量，也就是：布尔、整数、浮点数、复数或者字符串常量。
	//如果一个二元运算（非位移）的无类型操作数是不同类的，那么其结果是在如下列表中靠后显示的操作数的类：整数、 rune、浮点数、复数。

	//根据这段话，15 + 25 + 5.2 是常量表达式，因为这个表达式的操作数都是无类型的常量，因为其中有 5.2，它的默认类型是浮点型，
	//所以这个常量表达式的结果虽然是无类型的，但默认类型是浮点型。
}
