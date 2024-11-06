package train

import (
	"fmt"
	"reflect"
)

type Pe struct {
	age int
}

func (x Pe) IsStructureEmpty() bool {
	return reflect.DeepEqual(x, Pe{})
}

func Struct() {
	x := Pe{}

	if x.IsStructureEmpty() {
		fmt.Println("Structure is empty")
	} else {
		fmt.Println("Structure is not empty")
	}
}

/**
如何判断一个结构体为空

https://www.includehelp.com/golang/how-to-check-if-structure-is-empty.aspx
*/
