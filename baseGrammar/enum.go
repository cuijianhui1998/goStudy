package main

import "fmt"

//使用const和iota实现枚举
type Week int

const (
	Monday Week = iota
	Wenseday
)

func main() {

	fmt.Print(Monday, Wenseday)
}
