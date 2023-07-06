package train

import (
	"fmt"
	"strconv"
	"strings"
)

// Stringer fmt 包中定义的 Stringer 是最普遍的接口之一。
type Stringer interface {
	String() string
}

type IPAddr [4]byte

type Person struct {
	Name string
	Age  int
}

// 给 IPAddr 添加一个 "String() string" 方法
func (ip IPAddr) String() string {
	b := ip[:]
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ".")
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func StringerOut() {

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	//隐式自动调用
	fmt.Println(a.String(), z)
}
