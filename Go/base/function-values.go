package main

import (
	"fmt"
	"math"
)


func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main()  {
	v := func(x, y float64) float64 { // 接收函数的返回值
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(v(5, 1))

	fmt.Println(compute(v)) // 函数作为值传入另一个函数
	fmt.Println(compute(math.Pow))
}