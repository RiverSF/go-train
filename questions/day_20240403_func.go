package questions

/**

下面这段代码有什么缺陷：
func sum(x, y int)(total int, error) {
	return x+y, nil
}


解析：

在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。如果有多个返回值必须加上括号()；如果只有一个返回值且命名也必须加上括号()。
这里的第一个返回值有命名 total，第二个没有命名，所以错误。
*/
