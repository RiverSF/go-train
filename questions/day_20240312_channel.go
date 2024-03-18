package questions

import "sync"

func Q20240312() {
	var wg sync.WaitGroup
	foo := make(chan int)
	bar := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case foo <- <-bar:
		default:
			println("default")
		}
	}()
	wg.Wait()
}

func q20240312() {
	/**
	https://polarisxu.studygolang.com/posts/go/action/chained-channel-operations-in-a-single-select-case/

	对于 select 语句，在进入该语句时，会按源码的顺序对每一个 case 子句进行求值：这个求值只针对发送或接收操作的额外表达式。

	比如：

	// ch 是一个 chan int；
	// getVal() 返回 int
	// input 是 chan int
	// getch() 返回 chan int
	select {
	  case ch <- getVal():
	  case ch <- <-input:
	  case getch() <- 1:
	  case <- getch():
	}
	在没有选择某个具体 case 执行前，例子中的 getVal()、<-input 和 getch() 会执行
	这里有一个验证的例子：https://play.studygolang.com/p/DkpCq3aQ1TE。


	每次进入以下 select 语句时：

	select {
	case ch <- <-input1:
	case ch <- <-input2:
	}
	<-input1 和 <-input2 都会执行，相应的值是：A x 和 B x（其中 x 是 0-5）。但每次 select 只会选择其中一个 case 执行，所以 <-input1 和 <-input2 的结果，必然有一个被丢弃了，也就是不会被写入 ch 中。因此，一共只会输出 5 次，另外 5 次结果丢掉了。（你会发现，输出的 5 次结果中，x 比如是 0 1 2 3 4）

	而 main 中循环 10 次，只获得 5 次结果，所以输出 5 次后，报死锁。
	*/
}
