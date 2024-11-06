package train

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Interface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat 实现了 Abser
	describe(a)
	a = &v // a *Vertex 实现了 Abser
	describe(a)

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v

	fmt.Println(a.Abs())
}

func describe(i Abser) {
	fmt.Printf("(%v, %T)\n", i, i)

	/**
	在内部，接口值可以看做包含值和具体类型的元组：

	(value, type)

	接口值保存了一个具体底层类型的具体值。

	接口值调用方法时会执行其底层类型的同名方法
	*/
}
