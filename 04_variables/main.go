// 变量说明
// 类型
// int8 int16 int32 int64 uint8 uint16 uint32 uint64
// int uint
// float32 float64
// string bool
// byte(uint8的别名，用于区分字节值和8位无符号整数值)
// var 声明一个或多个变量
// 声明时可以不带变量类型，Go会自动推导
// 仅声明不初始化的情况下，会默认赋予"零值"
// := 是声明加初始化的简写形式
package  main

import "fmt"

func main() {
	a := 1
	b := "golang"
	c := 1.00
	d := true
	e := "hello"
	f := `hello`
	g := 'h' // 字符

	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", d, d)
	fmt.Printf("%v %T\n", e, e)
	fmt.Printf("%v %T\n", f, f)
	fmt.Printf("%v %T\n", g, g)

	// 先声明再赋值
	var h string
	h = "go"
	fmt.Println(h)

	// 一次声明多个变量
	var i, j int
	i, j = 1, 2
	fmt.Println(i, j)

	// 声明并赋值，可省略类型
	var k int = 1
	fmt.Println(k)

	var l, m = 1, 2
	fmt.Println(l, m)

	// 混合多种类型
	var n, o = 1, "name"
	fmt.Println(n, o)

	p := 1
	fmt.Println(p)

	q, r := 1, "name"
	fmt.Println(q, r)

	// var()形式
	var (
		s int // 默认0
		t string // 默认空字符串""
		u float64 // 默认0
		v bool // 默认false
		w = "foo"
		x = 1
	)
	fmt.Println(s, t, u, v, w, x)
}
