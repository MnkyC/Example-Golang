// 输出说明
// %v 原值
// %+v %v的基础上，展开结构体的字段名和值
// %#v Go语法形式的值
// %T Go语法形式的类型
// %d 十进制
// %f 浮点数
// %s 字符串
// %b 二进制
// %x 十六进制，字母形式小写
// %X 十六进制，字母形式大写
// %#x 十六进制，字母形式小写，添加前导
// %#X 十六进制，字母形式大写，添加前导
// %t 布尔值
// %q 单引号表示的字符字面值
// %p 指针
// %#p 指针，去掉前导
package main

import "fmt"

func main() {
	fmt.Println(100)
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println("2.0 / 1.0 = ", 2.0 / 1.0)
	fmt.Println("3.0 / 2.0 = ", 3.0 / 2.0)
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)

	fmt.Printf("%d, %b\n", 100, 100)
	fmt.Printf("%s \n", "hello")
	fmt.Printf("%d \t %b \t %#x \n", 100, 100, 100)
	fmt.Printf("%t\n", true)

	for i := 32; i < 127; i++ {
		fmt.Printf("%d \t %b \t %x \t %#x \t %q \t %p \n", i, i, i, i, i, &i)
	}
}