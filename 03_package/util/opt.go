package util

// OptName会被导出，因为大写字母开头
func OptName(name string) string {
	return optName(name)
}