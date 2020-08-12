// 包说明
// 每个go文件属于且仅属于一个包，文件名和包名可以不同
// 每个应用程序必须有且仅有一个名为main的包
// 包引入时的路径是根据$GOPATH/src这个环境变量为相对路径来进行引入
// 一般来说路径名最后一个字段和包名相同
// GOPATH存在多个就会逐一进行寻找，没有找到就会编译出错
// 包中大写字母开头的名字是外部可见的（PS: 汉字不区分大小写，所以不会导出)
// 包名也可以有别名，用于解决冲突
// 包中有多个go文件，Go构建工具会将其根据文件名(字母)排序，然后依次调用编译器编译
// 包的初始化是自下而上的，如a包导入了b包，那么在a包初始化的时候b包必须已经初始化过了
// 每个包只会被初始化一次，main包最后初始化
package main

import (
	"fmt"
	tool "github.com/MnkyC/Example-Golang/03_package/util"
	"github.com/MnkyC/Example-Golang/03_package/name"
)

func main() {
	fmt.Println(tool.OptName("go"))
	fmt.Println(tool.Name)
	fmt.Println(myname.Name) // 这里注意，要使用包名调用
}