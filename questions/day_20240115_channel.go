package questions

func Q20240115() {
	/**
	关于 channel，下面语法正确的是：

	A. var ch chan int
	B. ch := make(chan int)
	C. <- ch
	D. ch <-
	*/
}

var ch chan int

func q20240115() {
	/**
	A、B 都是声明 channel；
	C 读取 channel；
	写 channel 是必须带上值，所以 D 错误。
	*/
}
