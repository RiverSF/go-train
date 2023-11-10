package train

import (
	"fmt"
	"sort"
)

// SortMap map无序转有序
func SortMap() {
	peoples := map[int]string{
		1: "好孩子",
		6: "老人家",
		5: "好大夫",
	}

	var keys []int
	for i, _ := range peoples {
		keys = append(keys, i)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%v:%v;", k, peoples[k]) //1:好孩子;5:好大夫;6:老人家
	}
}
