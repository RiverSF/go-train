package questions

/**
下面赋值正确的是：

A. var x = nil
B. var x interface{} = nil
C. var x string = nil
D. var x error = nil
*/

/**
参考答案及解析：BD。这道题考的知识点是 nil。nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量。强调下 D 选项的 error 类型，它是一种内置接口类型，看它的源码就知道，所以 D 是对的。

type error interface {
	Error() string
}
*/
