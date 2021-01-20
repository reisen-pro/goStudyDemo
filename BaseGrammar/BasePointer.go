package main

import "fmt"

// 一个指针变量通常缩写为 ptr。
var ptr *int

func main() {
	var a int = 10
	// 一个指针变量指向了一个值的内存地址。
	// 类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：
	fmt.Printf("变量的地址: %x\n", &a)

	// var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明：
	// var var_name *var-type
	var ip *int     /* 指向整型*/
	var fp *float32 /* 指向浮点型 */
	fmt.Println(ip, fp)
	// 空指针的判断
	fmt.Println(ip == nil)

	// 讲a的指针地址赋给ip
	ip = &a
	// 打印ip这个指针的地址的值
	fmt.Println("ip的值", *ip)

}
