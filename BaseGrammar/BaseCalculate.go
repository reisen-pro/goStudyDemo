package main

import "fmt"

func main() {
	var a = 21
	var b = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	// a = a++ // 这是不允许的，会出现变异错误 syntax error: unexpected ++ at end of statement
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)

	var aa int = 4
	var bb int32
	var cc float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - aa 变量类型为 = %T\n", aa)
	fmt.Printf("第 2 行 - bb 变量类型为 = %T\n", bb)
	fmt.Printf("第 3 行 - cc 变量类型为 = %T\n", cc)

	/*  & 和 * 运算符实例 */
	ptr = &aa /* 'ptr' 包含了 'aa' 变量的地址 */
	fmt.Printf("aa 的值为  %d\n", aa)
	fmt.Printf("*ptr 为 %d\n", *ptr)
	fmt.Printf("ptr 为 %d\n", ptr)

	for true {
		fmt.Printf("这是无限循环。\n")
	}
}
