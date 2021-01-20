package main

import "fmt"

// 全局变量是允许声明但不使用的 同一类型的多个变量可以声明在同一行
var (
	width            int32 = 40
	height           int32 = 180
	int1, int2, int3       = 1, 2, 3
)

// 定义常量 必须赋初始值
const world string = "world"

// 隐式
const hello = "hello"

// 常量还可以用作枚举
const (
	Female  = 1
	Male    = 2
	Unknown = 0
)

// iota，特殊常量，可以认为是一个可以被编译器修改的常量。 const 中每新增一行常量声明将使 iota 计数一次
// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 b和c可以简写 = iota
const (
	a = iota
	b = iota
	c = iota
)

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {
	// 这种因式分解关键字的写法一般用于声明全局变量 定义在全局
	fmt.Println("width", width, "height", height)

	fmt.Println(i, j, k, l)
	fmt.Println(a, b, c)

	x := 2
	y := 2
	// ^ 亦或运算符
	fmt.Println(x ^ y)
}
