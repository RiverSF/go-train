package questions

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3},
}

func Q20231120() {
	//m["foo"].x = 4
	//fmt.Println(m, m["foo"].x)

	//正确更新方式
	//mm := m["foo"]
	//mm.x = 4
	//m["foo"] = mm
	//fmt.Println(m["foo"].x)
}

func q20231120() {
	//在Go语言中，要更新map的值，首先需要使用键来获取对应的值，然后对值进行更新操作。
	//由于map的值是按值传递的，因此需要将结构体的副本赋值给一个新的变量，以便在更新时不会影响原始结构体。
	//如果直接修改map中的值，原始结构体将不会被更改。

	//如果map的值是指针类型，可直接修改，如下
	/**

	people := make(map[string]*Person)

	p1 := &Person{Name: "Alice", Age: 25}
	people["alice"] = p1

	//更新p1的年龄
	people["alice"].Age += 1

	*/
}
