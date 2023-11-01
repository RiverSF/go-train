package questions

import "fmt"

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
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
	f1()、f2()、f3() 函数分别返回什么？
	*/
	fmt.Println(f1(), f2(), f3())
}

func q20231030() {
	// defer 在跳出当前方法时执行; 先进后出执行
	// defer 函数【不带参数时】：匿名函数在 return【之前】执行
	// defer 函数【带参数时】： 匿名函数在 return【之后】执行

	// f1
	// 由于 defer 语句在函数返回之前执行，因此当 return 0 被执行时，defer 语句也会被执行

	// f2
	// 返回变量为 r，defer 方法中变量是 t，因此 r = 5

	// f3
	// defer 语句注册的匿名函数带有参数，因此它会在函数返回后执行该匿名函数
}
