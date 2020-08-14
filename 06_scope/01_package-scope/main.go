// 包作用域说明
package main

import "fmt"

var x = 100 // 全局

func main() {
	fmt.Println(x) // 100
	foo()
}

func foo() {
	fmt.Println(x) // 100
}