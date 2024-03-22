package questions

import "fmt"

func Q20240322() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}

func q20240322() {
	/**
	map 是无序的，如果第一次循环到 A，则输出 3；否则输出 2。

	map 本身是引用类型，range 拷贝的副本删除元素会直接修改源map
	*/
}
