package main

import "fmt"

func main() {

}

// 无参数
func foo() int {
	return 100
}
// 多个参数
func add(x int, y int) int {
	return x+y
}
// 无返回
func printSome(n int) {
	fmt.Println(n)
}
// 多值返回
func swap(x, y string) (string, string) {
	return y, x
}
// 命名返回值:应当仅用在短函数中。在长的函数中它们会影响代码的可读性。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}