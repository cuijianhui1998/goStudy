package main

import "fmt"

//类型别名

func main() {
	//给int取一个别名
	type NewInt = int
	//将a定义为int类型
	type a int

	var b NewInt
	fmt.Print(b) //b的值为int型的nil值

	var c a
	fmt.Print(c) //c也为int型的nil值,有点感觉c是int的子类(只是猜测)

}
