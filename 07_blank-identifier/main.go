// 空白标识符说明
// Go中声明的变量没有使用会报错
// 我们可以用空白标识符(下划线)来进行舍弃
package main

import "fmt"

func main() {
	a := "a"
	//b := "b" // 报错 b declared but not used
	//_ := "c" // 报错，因为:=左边必须要有值
	_ = "c"
	c, _ := "c", "d"
	fmt.Println("a - ", a)
	fmt.Println("c - ", c)
}