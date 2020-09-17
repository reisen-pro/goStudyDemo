package main

import "fmt"

func main() {
	fmt.Println(max(10, 20))
	x, y := swap("world", "hello")
	fmt.Println(x, y)
}

func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}
