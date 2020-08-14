// 遍历说明
// Go唯一的循环结构
// 支持continue和break

package main

import "fmt"

func main() {
	// 实现while
	i := 1
	for i < 3 {
		fmt.Println(i)
		i = i + 1 // i++
	}

	// 常规
	for j := 1; j < 3; j++ {
		fmt.Println(j)
	}

	// 嵌套
	for i := 1; i < 3; i++ {
		for j := 1; j < 3; j++ {
			fmt.Println(i, j)
		}
	}

	// 无条件
	for {
		fmt.Println("go")
		break
	}

	for {
		i++
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
		if i > 10 {
			break
		}
	}

	for i := 0; i < 5; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}