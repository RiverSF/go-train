package questions

import "fmt"

type Stu struct{}

func (p *Stu) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *Stu) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	Stu
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func Q20240130() {
	/**
	下面这段代码输出什么？
	*/

	t := Teacher{}
	t.ShowB()
}

func q20240130() {
	/**
	参考答案及解析：teacher showB。

	知识点：结构体嵌套。
	在嵌套结构体中，Stu 称为内部类型，Teacher 称为外部类型；通过嵌套，内部类型的属性、方法，可以为外部类型所有，就好像是外部类型自己的一样。
	此外，外部类型还可以定义自己的属性和方法，甚至可以定义与内部相同的方法，这样内部类型的方法就会被“屏蔽”。这个例子中的 ShowB() 就是同名方法。
	*/
}
