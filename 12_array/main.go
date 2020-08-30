// 数组说明
// 数组：具有特定长度的元素的编号序列
// 元素的类型和长度都是数组类型的一部分，[3]int和[4]int是不同的类型
// 数组属于值类型，进行拷贝和修改，都不会影响原值
// 索引从0开始，越界会引发panic
// 数组定义后，每个位置会有默认值，默认"零值"
// 利用下标可以设置和获取一个值
// 内置函数len用于获取数组长度
// 支持多维数组
// 数组可以进行比较，所有元素都相等才相等

package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty a:", a)
	fmt.Println("empty len:", len(a))
	fmt.Println(a[0], a[4])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	var x [58]string
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	// 遍历
	for _, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, []byte(v))
	}
	fmt.Println(x)
	fmt.Println(x[42])

	var y [100]int
	fmt.Println(len(y))
	fmt.Println(y[10])
	for i := 0; i < 100; i++ {
		y[i] = i
	}
	// 遍历
	for i, v := range y {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 20 {
			break
		}
	}

	var z [100]byte
	fmt.Println(len(z))
	fmt.Println(z[10])
	for i := 0; i < 100; i++ {
		z[i] = byte(i)
	}
	for i, v := range z {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 20 {
			break
		}
	}

	// 多维
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// 利用地址修改值
	arr := [3]int{1, 2, 3}
	fmt.Println("before:", arr)
	changeVal(&arr)
	fmt.Println("after:", arr)
}

func changeVal(v *[3]int) {
	(*v)[0] = 100
}