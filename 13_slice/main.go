// 切片说明
// 提供了比数组更强大的序列接口
// 用 make 来创建切片
// 和数组一样，可以利用下标设置和获取一个值
// 内置函数len用于获取预期的长度
// 增加了内置函数append，用于追加元素，返回新的切片值（注意接收返回值）
// 增加了内置函数copy，可以复制切片
// 增加了切片运算符 [low:high]，返回新切片（注意，不包括high位置的值）
// 可以组成多维数据结构，和多维数组不同，内部切片的长度可以变化

package main

import "fmt"

func main() {
	var a []int // 声明
	fmt.Println("a:", a, "a == nil:", a == nil) // [] true
	a = make([]int, 3)
	fmt.Println("empty a:", a)
	fmt.Println("empty len:", len(a))
	fmt.Println("empty cap:", cap(a))
	fmt.Printf("%T\n", a)
	fmt.Println(a[0], a[1])

	a[0] = 100
	a[2] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[2])
	fmt.Println("len:", len(a))
	fmt.Println("cap:", cap(a))

	// append后会增大length，和原capacity比较，不够就进行扩容（增大capacity），扩容规则：原capacity * 2
	a = append(a, 10)
	fmt.Println("len:", len(a), "capacity:", cap(a)) // 4, 6
	a = append(a, 1, 2)
	fmt.Println("len:", len(a), "capacity:", cap(a)) // 6, 6
	fmt.Println("apd:", a)
	a = append(a, 11)
	fmt.Println("len:", len(a), "capacity:", cap(a)) // 7, 12
	a = append(a, 1, 2)
	fmt.Println("len:", len(a), "capacity:", cap(a)) // 9, 12
	fmt.Println("apd:", a)

	// 指定length和capacity，注意，capacity >= length
	x := make([]string, 3, 5)
	x[0] = "go"
	x[1] = "lang"
	x[2] = "golang"
	fmt.Println(x[2])

	// 复制
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println("cpy:", b)

	l := a[2:5]
	fmt.Println("a[2:5]:", l)

	l = a[:5]
	fmt.Println("a[:5]:", l)

	l = a[2:]
	fmt.Println("a[2:]:", l)

	l = a[:]
	fmt.Println("a[:]:", l)

	// 获取字符对应的数值
	fmt.Println("golang"[2])

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 遍历
	s := []int{1, 3, 5, 7, 9}
	for i, v := range s {
		fmt.Println(i, "-", v)
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(i, "-", s[i])
	}

	// 添加切片
	s1 := []int{10, 11, 12}
	s = append(s, s1...)
	fmt.Println("append s1:", s)

	// 删除切片元素
	s = append(s[:2], s[3:]...)
	fmt.Println("delete s:", s)

	// 多维
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // [[0] [1 2] [2 3 4]]

	// 赋值
	name := []string{}
	names := [][]string{}
	fmt.Println(name) // []
	fmt.Println(names) // []
	fmt.Println(name == nil) // false
	fmt.Println(names == nil) // false

	var name1 []string
	var names1 [][]string
	fmt.Println(name1) // []
	fmt.Println(names1) // []
	fmt.Println(name1 == nil) // true
	fmt.Println(names1 == nil) // true

	name2 := make([]string, 4)
	names2 := make([][]string, 4)
	fmt.Println(name2) // [   ]
	fmt.Println(names2) // [[] [] [] []]
	fmt.Println(name2 == nil) // false
	fmt.Println(names2 == nil) // false

	//name[0] = "golang" // 引发panic异常，index out of range [0] with length 0
	name = append(name, "golang") // [golang]
	fmt.Println(name)
	fmt.Println(name[0]) // golang

	//name1[0] = "golang" // 引发panic异常，index out of range [0] with length 0
	name1 = append(name1, "golang") // [golang]
	fmt.Println(name1)
	fmt.Println(name1[0]) // golang

	name2[0] = "golang" // [golang   ]
	name2 = append(name2, "golang") // [golang    golang]
	fmt.Println("len:", len(name2), "capacity:", cap(name2)) // 5, 8
	fmt.Println(name2)
	fmt.Println(name2[0]) // golang
	fmt.Println(name2[1]) // 空字符串
	fmt.Println(name2[2]) // 空字符串
	fmt.Println(name2[3]) // 空字符串
	fmt.Println(name2[4]) // golang

	// 嵌套
	var records [][]string
	name3 := make([]string, 4)
	name3[0] = "golang"
	name3[1] = "ts"
	name3[2] = "100.00"
	name3[3] = "100.00"
	//records[0] = name3 // // 引发panic异常，index out of range [0] with length 0
	records = append(records, name3)
	name4 := make([]string, 4)
	name4[0] = "java"
	name4[1] = "python"
	name4[2] = "100.00"
	name4[3] = "96.00"
	records = append(records, name4)
	fmt.Println(records) // [[golang ts 100.00 100.00] [java python 100.00 96.00]]

	names5 := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		name5 := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			name5 = append(name5, j)
		}
		names5 = append(names5, name5)
	}
	fmt.Println(names5) // [[0 1 2 3] [0 1 2 3] [0 1 2 3]]
	fmt.Println(names5[1]) // [0 1 2 3]

	name6 := make([]int, 1)
	fmt.Println(name6[0]) // 0
	name6[0] = 1
	fmt.Println(name6[0]) // 1
	name6[0]++
	fmt.Println(name6[0]) // 2
}