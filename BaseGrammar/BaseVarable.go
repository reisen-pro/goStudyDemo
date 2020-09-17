package main

import (
	"fmt"
)

func main() {

	/* 定义变量 */

	// 第一种
	// 定义两个变量  var 变量名
	// 不需要加;进行 行之间分格
	var helloStr = "hello "
	var goStr = "go"

	fmt.Println(helloStr + goStr)

	// 第二种
	// 可以根据值的内容自行判断类型
	var a = 12
	var b = 27.0
	var c = true
	fmt.Println(a, b, c)

	// 第三种
	// :=  等价于 var name string = "name"
	// := 左侧的变量不应该是已经被声明过的，否则会导致编译错误
	// 这种不带声明格式的只能在函数体中出现 fun main(){ name:="name" }
	name := "name"
	fmt.Println("name", name)

	// 可以一次声明多个变量
	var str3, str4 = "hello", "world"

	fmt.Println(str3 + " " + str4)

	// 声明了变量和类型 不带初始值的 会有一个默认值 string 类型 默认""
	var str string
	fmt.Println("str: ", str)

	/* 定义数字类型 */

	// 有符号 8 位整型 (-128 到 127)
	var int8 int8 = 127
	// 有符号 16 位整型 (-32768 到 32767)
	var int16 int16 = 5500
	// 有符号 32 位整型 (-2147483648 到 2147483647)
	var int32 int32 = 55555
	// 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
	var int64 int64 = 5555555555
	fmt.Println("int8", int8)
	fmt.Println("int16", int16)
	fmt.Println("int32", int32)
	fmt.Println("int64", int64)

	// IEEE-754 32位浮点型数
	var float32 float32
	// IEEE-754 64位浮点型数
	var float64 float64
	fmt.Println("float32", float32)
	fmt.Println("float64", float64)

	// 布尔类型 默认为false
	var boolean bool
	fmt.Println("bool", boolean)

	// _表示抛弃值
	_, numb, str := numbers() //只获取函数返回值的后两个
	fmt.Println(numb, str)
}

//一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
