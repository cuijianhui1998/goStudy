package main

import "fmt"

//++不再是一个操作符,而是一条语句

func main() {
	a := 1
	//不支持在赋值时自增
	//b := a++
	fmt.Println(a)
	a++
	//不支持前置的自增
	//++a
	b := a
	fmt.Println(b)
}
