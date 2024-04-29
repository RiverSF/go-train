package questions

import (
	"fmt"
)

type People1 interface {
	Show()
}

type Student1 struct{}

func (stu *Student1) Show() {

}

func live() People1 {
	var stu *Student1
	return stu
}

func Q20240429() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

/**
分析：

我们需要了解interface的内部结构，才能理解这个题目的含义。（源码基于 Go1.17）

interface 在使用的过程中，共有两种表现形式

一种为空接口(empty interface)，定义如下：

var MyInterface interface{}

另一种为非空接口(non-empty interface), 定义如下：

type MyInterface interface {
	function()
}
这两种 interface 类型在底层分别用两种struct表示，空接口为eface, 非空接口为iface。

空接口 eface
空接口 eface 结构，由两个属性构成，一个是类型信息 _type，一个是数据信息。其数据结构声明如下：

type eface struct {      // 空接口
    _type *_type         // 类型信息
    data  unsafe.Pointer // 指向数据的指针(go 语言中特殊的指针类型 unsafe.Pointer 类似于 c 语言中的void*)
}

非空接口 iface
iface 表示 non-empty interface 的数据结构，非空接口初始化的过程就是初始化一个 iface 类型的结构，其中data的作用与 eface 的相同，这里不再多加描述。

type iface struct {
  tab  *itab
  data unsafe.Pointer
}

stu 是一个指向 nil 的空指针，但是最后return stu 会触发匿名变量 People1 = stu 值拷贝动作，所以最后live()放回给上层的是一个People1 interface{}类型，也就是一个iface struct{}类型。
stu 为 nil，只是iface中的 data 为 nil 而已。 但是iface struct{}本身并不为 nil.
*/
