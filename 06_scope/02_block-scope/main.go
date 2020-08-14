// 块作用域说明
package main

import "fmt"

var i = 0 // 全局

// 闭包
func increment() int {
	i++
	return i
}

// 闭包
func wrapper() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	x := 100 // 局部
	fmt.Println(x) // 100
	foo()

	{
		fmt.Println(x)
		y := "go"
		fmt.Println(y)
	}
	//fmt.Println(y) // 报错 undefined: y

	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2

	// 闭包
	y := 0
	increment1 := func() int {
		y++
		return y
	}
	fmt.Println(increment1()) // 1
	fmt.Println(increment1()) // 2

	increment2 := wrapper()
	fmt.Println(increment2()) // 1
	fmt.Println(increment2()) // 2
}

func foo() {
	//fmt.Println(x) // 报错 undefined: x
}