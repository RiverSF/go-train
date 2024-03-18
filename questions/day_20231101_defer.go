package questions

import "fmt"

func f() {
	defer fmt.Println("D")
	fmt.Println("F")
}

func Q20231102() {
	/**
	下面这段代码正确的输出是什么？
	*/

	f()
	fmt.Println("M")
}
