// fmt_test.go
package main

import (
	"fmt"
)

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

type T struct {
	a int
	b float32
	c string
}

var t = &T{1, 1.2, "abc\tdef"}

func main() {
	fmt.Printf("No01: %v\n", t)
	fmt.Printf("No02: %+v\n", t)
	fmt.Printf("No03: %#v\n", t)
	fmt.Printf("No04: %#v\n", timeZone)
	fmt.Printf("No05: %T\n", t)
	//fmt.Printf("No06: %%\n", t)

	fmt.Printf("No07: %t\n", 1 == 1)
	fmt.Printf("No08: %b\n", 5)
	fmt.Printf("No09: %o\n", 12)
	fmt.Printf("No10: %x\n", 12)
	fmt.Printf("No10: %#x\n", 12)
	fmt.Printf("No11: %X\n", 12)
	fmt.Printf("No12: %d\n", 5)
	fmt.Printf("No13: %c\n", 'a')
	fmt.Printf("No14: %q\n", 'a')
	fmt.Printf("No15: %U\n", 5)

	fmt.Printf("No16: %e\n", 10.2)
	fmt.Printf("No17: %E\n", 10.2)
	fmt.Printf("No18: %G\n", 10.20+2i)
	fmt.Printf("No19: %f\n", 1.22)
	fmt.Printf("No20: %g\n", 1.222)

	fmt.Printf("No21: %s\n", "Go语言")
	fmt.Printf("No22: %q\n", "Go语言")
	fmt.Printf("No23: %x\n", "golang")
	fmt.Printf("No24: %X\n", "golang")

	fmt.Printf("No25: %p\n", t)

	fmt.Printf("No26: %.*d\n", 10, 123)
	fmt.Printf("N027: %6.*f\n", 2, 888.666)
}
