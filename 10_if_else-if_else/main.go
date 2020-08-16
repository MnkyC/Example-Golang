// 条件语句说明
// 可以只有if，没有else
// 注意，不要用括号包围条件语句，其次，必须使用花括号
// Go没有三元组if，必须要用完整的if语句
package main

import "fmt"

func main() {
	if true {
		fmt.Println("true")
	}

	if false {
		fmt.Println("false")
	}

	if !true {
		fmt.Println("false")
	}

	if !false {
		fmt.Println("true")
	}

	a := true
	if b := "go"; a {
		fmt.Println(b)
	}

	c := true
	if d := "golang"; c {
		fmt.Println(d)
	}
	//fmt.Println(d) // 报错 undefined: d

	if true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if false {
		fmt.Println("go")
	} else if true {
		fmt.Println("golang")
	} else {
		fmt.Println("gogogo")
	}

	if false {
		fmt.Println("go")
	} else if false {
		fmt.Println("gol")
	} else if true {
		fmt.Println("golang")
	} else {
		fmt.Println("gogogo")
	}

	if 4 % 2 == 0 {
		fmt.Println("4 is divisible by 2")
	}

	if 3 % 2 == 0 {
		fmt.Println("3 is even")
	} else {
		fmt.Println("3 is odd")
	}

	if x := 9; x < 0 {
		fmt.Println(x, "is negative")
	} else if x < 10 {
		fmt.Println(x, "has 1 digit")
	} else {
		fmt.Println(x, "has multiple digits")
	}
}