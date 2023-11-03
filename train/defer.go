package train

func Defer() {
	/**
	defer 详解：参考 https://blog.csdn.net/lxw1844912514/article/details/130023459
	*/

	// 一、defer 执行顺序
	/**
	1. 当go执行到一个defer后的语句时，会将defer后的语句先压入到一个栈中，然后继续执行函数下一个语句

	2. 在defer将语句放入到栈中时，如果有参数传递时，会将相关参数的值拷贝同时入栈，不影响上层变量结果

	3. 当函数其他命令执行完毕后，在从defer栈中，依次从栈顶去除语句执行
	(注: 遵守栈 先进后出的机制来输出)
	*/

	// 二、defer 触发时机
	/**
	1. 包裹着 defer 语句的函数返回时
	2. 包裹着 defer 语句的函数执行到最后
	3. 当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行panic
	*/

	// 三、defer，return，返回值  的执行顺序
	/**
	1. 先给返回值赋值
	2. 执行defer语句
	3. 包裹函数return返回

	示例代码见 每日一题 day_20231030.go
	*/

}
