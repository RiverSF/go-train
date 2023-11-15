package train

import "fmt"

//func (i int) PrintInt ()  {
//	fmt.Println(i)
//}
//
//func Func() {
//	var i int = 1
//	i.PrintInt()
//}

/**
上述代码会报编译错误。
原因解析：
	上面的代码基于 int 类型创建了 PrintInt() 方法，由于 int 类型和方法 PrintInt() 定义在不同的包内，所以编译出错。
	解决的办法可以定义一种新的类型
*/

type Myint int

func (i Myint) PrintInt() {
	fmt.Println(i)
}
func Func2() {
	var i Myint = 1
	i.PrintInt()
}

// 方法 总结
/**
1.  你只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法
2.	不能为内建类型声明方法。
*/
