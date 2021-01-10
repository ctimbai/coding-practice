package main

import "fmt"

type Vertex struct {
	x, y float64
}

var gm = map[string]Vertex{
	"hello": {
		2.3, 3.4,
	},
	"world": {
		4.5, 8.3,
	},
}

func main()  {
	// map 的初始化
	var m map[int]string
	fmt.Println(m)
	m = map[int]string{1:"hello"}
	fmt.Println(m)
	fmt.Println(m[1])

	m1 := map[int]string{1:"world"}
	fmt.Println(m1)

	m2 := make(map[int]string)
	m2[1] = "goworld"
	fmt.Println(m2)

	fmt.Println(gm)

	// map 的增删改查
	m3 := make(map[string]int)
	m3["hello"] = 12
	m3["world"] = 34
	fmt.Println(m3)

	delete(m3, "hello")
	fmt.Println(m3)

	v, ok := m3["world"]
	if ok == true {
		fmt.Println(v)
	} else {
		fmt.Println("nil")
	}
}
