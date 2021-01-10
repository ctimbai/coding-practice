package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 { // 方法就是带接受者参数的正常函数
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func Abs1(v Vertex) float64 { // 普通函数
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

// 指针接受者
func (v *Vertex) scale(f float64) {
	v.x *= f
	v.y *= f
}

func main()  {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.scale(10) // 隐式转换
	fmt.Println(v.Abs())

	fmt.Println(Abs1(v))
}