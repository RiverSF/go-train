package train

import "golang.org/x/tour/tree"
import "fmt"
import "strings"
import "strconv"

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	l := t.Left
	r := t.Right
	v := t.Value
	//fmt.Println(l, p, r)

	if l != nil {
		Walk(l, ch)
	}

	ch <- v

	if r != nil {
		Walk(r, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {

	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var x, y []string

	for i := 0; i < 10; i++ {
		x = append(x, strconv.Itoa(<-ch1))
		y = append(y, strconv.Itoa(<-ch2))
	}
	xx := strings.Join(x, ",")
	yy := strings.Join(y, ",")
	fmt.Println(xx, yy)

	return xx == yy
}

func ChannelOut() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
