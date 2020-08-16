// switch说明
// case语句中可以使用逗号分割多个表达式
// fallthrough 强制执行后面的case代码，但是不能用在最后一个分支
// 可以用来代替没有表达式的if语句
package main

import "fmt"

type myname struct {
	name string
}

func SwitchOnType(x interface{}) {
	switch x.(type) { // 断言，判断x是否是这种类型
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case myname:
		fmt.Println("myname")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	switch "Four" {
	case "One":
		fmt.Println("One")
	case "Two":
		fmt.Println("Two")
	case "Three":
		fmt.Println("Three")
	default:
		fmt.Println("Check")
	}

	// fallthrough
	switch "Three" {
	case "One":
		fmt.Println("One")
	case "Two":
		fmt.Println("Two")
	case "Three":
		fmt.Println("Three")
		fallthrough
	case "Four":
		fmt.Println("Four")
		fallthrough
	case "Five":
		fmt.Println("Five")
	case "Six":
		fmt.Println("Six")
		//fallthrough // 报错 Cannot fallthrough final case in switch
	}

	// 逗号分隔多个表达式
	switch "Four" {
	case "One", "Four":
		fmt.Println("One, Two")
	case "Two", "Three":
		fmt.Println("start with T")
	case "Five", "Six":
		fmt.Println("Five / Six")
	}

	// 使用表达式
	Num := "ZERO"

	switch {
	case len(Num) == 2:
		fmt.Println("Expression 'len(Num) == 2' is always 'false'")
	case Num == "One":
		fmt.Println("Expression 'Num == \"One\"' is always 'false'")
	case Num == "Two":
		fmt.Println("Expression 'Num == \"Two\"' is always 'false'")
	case Num == "Three", Num == "Four":
		fmt.Println("Three / Four")
	case Num == "Five":
		fmt.Println("Expression 'Num == \"Five\"' is always 'false'")
	case Num == "Six":
		fmt.Println("Expression 'Num == \"Six\"' is always 'false'")
	default:
		fmt.Println("default")
	}

	switch {
	case 1 < 2:
		fmt.Println("smaller")
	default:
		fmt.Println("bigger")
	}

	// 判断类型
	SwitchOnType(1)
	SwitchOnType("golang")
	var t = myname{"golang"}
	SwitchOnType(t)
	SwitchOnType(t.name)

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("golang")
}
