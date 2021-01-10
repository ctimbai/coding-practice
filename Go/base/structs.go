package main

import "fmt"

type Point struct {
	X, Y int
}

func main()  {

	// 结构体初始化
	var p1 Point = Point{1,2}
	p2 := Point{2, 3}
	fmt.Println(p1, p2)

	// 写成列表的形式
	var (
		v1 = Point{} // 默认 X = 0, Y = 0
		v2 = Point{X:1} // Y:0 隐式赋值
		v3 = Point{X: 1, Y: 3} // 手动以 name:xx 的形式赋值
		v4 = &Point{1, 5}
	)
	fmt.Println(v1, v2, v3, v4)

	// 结构体赋值
	p3 := p1
	fmt.Println(p3, p1)

	p3.X = 10
	fmt.Println(p3.X, p1.X)

	// 结构体指针
	p4 := &p1
	fmt.Println(*p4)

	p4.X = 10 // 这种写法是隐式间接引用，写成 (*p4).X = 10 也可以
	p4.Y = 1.3e5
	fmt.Println(*p4, p1)
}