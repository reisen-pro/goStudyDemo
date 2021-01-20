package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	price   float32
}

func main() {
	// 直接构造
	fmt.Println(Book{"Go 语言", "Rob Pike", "Go语言学习", 99.9})
	// 指定key-value形式
	fmt.Println(Book{title: "Go 语言", author: "Rob Pike", subject: "Go语言学习", price: 99.9})
	// 忽略的字段为默认值
	fmt.Println(Book{title: "Go 语言", author: "Rob Pike"})

	var java Book
	var php Book
	java.title = "java"
	java.author = "James Gosling"
	java.subject = "java study"
	java.price = 99.9

	php.title = "php"
	php.author = "Rasmus Lerdorf"
	php.subject = "php study"
	php.price = 99.9

	/* 打印 Book1 信息 */
	fmt.Printf("java title : %s\n", java.title)
	fmt.Printf("java author : %s\n", java.author)
	fmt.Printf("java subject : %s\n", java.subject)
	fmt.Printf("java price : %d\n", java.price)

	/* 打印 Book2 信息 */
	fmt.Printf("Book title : %s\n", php.title)
	fmt.Printf("Book author : %s\n", php.author)
	fmt.Printf("Book subject : %s\n", php.subject)
	fmt.Printf("Book price : %d\n", php.price)
}
