package util

// optName不会导出，但可以被同一个包内的其他函数直接调用
func optName(name string) string {
	return name + "lang"
}
