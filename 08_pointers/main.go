// 指针说明
// 可以通过传递内存地址，通过地址修改存储的值
// 这种传递方式使得程序更高效，不必传递大量数据，只需要传递地址即可
// 指针是一个值
// &v 获取v的内存地址，即指向v的指针
package main

import "fmt"

func main() {
	a := 100
	fmt.Println("a's memory address:", &a) // 内存地址 0xc00001a088
	fmt.Printf("%d \n", &a) // 824633827464

	// 使用内存地址
	var b float64
	fmt.Print("Enter:")
	fmt.Scan(&b)
	c := b * 1.01
	fmt.Println(c)

	// d的类型是int指针，指向存储100的内存地址的指针
	var d = &a
	fmt.Println("d's memory address:", d) // 0xc00001a088

	// 取消引用，在这里*表示运算符
	fmt.Println("*d:", *d) // 100

	// 使用指针
	*d = 1        // 地址对应的值改为1
	fmt.Println("修改指针d所指向的内存地址中存储的值为", *d) // 1
	fmt.Println("d's memory address:", d) // 0xc00001a088
	fmt.Println("a:", a) // 1

	// 不用指针进行传递
	e := 1
	fmt.Printf("e's memory address in main:%p\n", &e) // 0xc00001a0c0
	fmt.Println(&e) // 0xc00001a0c0
	foo(e)
	fmt.Println("值传递后原值不变:", e) // e仍然是1

	// 用指针传递
	f := 1
	fmt.Printf("f's memory address in main:%p\n", &f) // 0xc00001a0d8
	fmt.Println(&f) // 0xc00001a0d8
	foo1(&f)
	fmt.Println("指针传递后原值改变:", f) // g变为0
}

func foo(v int) {
	fmt.Printf("memory address in foo:%p\n", &v) // 0xc00001a0c8
	fmt.Println(&v) // 0xc00001a0c8
	v = 0
}

func foo1(v *int) {
	fmt.Printf("memory address in foo1:%p\n", v) // 0xc00001a0d8
	fmt.Println(v) // 0xc00001a0d8
	*v = 0
}