// 常量说明
// Go支持字符，字符串，布尔和数字的常量
// const 声明一个或多个常量
// 声明时没有显式指定类型，Go会自动推导
// 常量表达式会以任意精度进行算术计算
// 常量之间算术运算，逻辑运算和比较运算的结果都是常量
// 常量的运算都在编译期间完成
// 常量可以用于实现枚举
package main

import (
	"fmt"
	"math"
)

const s string = "hello"

const (
	a = 1
	b = "go"
)

// 借助iota这个特殊类型实现枚举
// 第一个值必须写iota，后面的可写可不写
const (
	c = iota // 0
	d =	iota // 1
	e = iota // 2
)

const (
	f = iota // 0
	g 		 // 1
	h 		 // 2
)

// 用下划线来舍弃
const (
	_ = iota       // 0
	i = iota * 100 // 1 * 100
	j = iota * 100 // 2 * 100
)

// 位操作
const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

// 省略初始化表达式会自动使用前面常量的初始化表达式，第一个不能省略
const (
	k = 1
	l // 1
	m = 2
	n // 2
)

func main() {
	fmt.Println(s)

	const x = 100
	const y = 1e20 / x
	fmt.Println(y)
	fmt.Println(int64(y))
	fmt.Println(math.Sin(y))

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

	fmt.Println("binary\t\t\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t\t\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)

	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
}