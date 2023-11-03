package questions

import "fmt"

func f0() int { //匿名返回值

	r := 1
	defer func() {
		r++
	}()

	return r
}

func f1() (r int) { //有名返回值

	defer func() {
		r++
	}()

	return 1
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {

	defer func(r int) {
		r = r + 5
	}(r)

	return 1
}

func Q20231030() {
	/**
	f0()、f1()、f2()、f3() 函数分别返回什么？
	*/
	fmt.Println(f0(), f1(), f2(), f3())
}

func q20231030() {
	// f0 = 1
	// f0 是匿名返回值，匿名返回值是在 return 执行时被声明，因此当 defer 声明时，还不能访问到匿名返回值，defer 的修改不会影响到匿名返回值。

	// f1 = 2
	// f1 是有名返回值，先给 r 赋值，r=1，执行 defer 语句，r=2，然后 return

	// f2 = 5
	// 返回变量为 r，defer 方法中变量是 t，因此 r = 5

	// f3 = 1
	// defer 函数有参数传递，在defer入栈时就会对参数进行拷贝传递，因此参数作用域仅限defer函数内部，不会影响外层return结果
}
