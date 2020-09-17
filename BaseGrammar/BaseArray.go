package main

import "fmt"

func main() {

	// 定义一个数组 var 数组名 = [个数]类型{赋值}
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)

	// 定义一个数组并循环打印 [...] 会根据元素的个数来设置数组的大小
	var array = [...]int{1, 2, 3, 4, 5}
	for i := range array {
		fmt.Println(i)
	}

	var arr [10]int
	for arrInt := 0; arrInt < len(arr); arrInt++ {
		arr[arrInt] = arrInt + 100
	}
	for num := range arr {
		fmt.Println(num)
	}
}
