package main

import "fmt"

//变量作用域

//定义了一个全局变量
var global = 4

func local(a int) int {
	//参数,返回值,在函数中定义的变量都是局部变量
	var b = 3
	//函数体中并未声明global变量,但是仍然可以使用global变量
	fmt.Println(global)
	//声明一个局部变量global,与全局变量重名,程序会由内向外查询
	var global int
	fmt.Println(global)
	return b
}

func main() {
	//在main中定义一个变量名为b的变量,不会和local()中的b冲突
	b := local(1)
	fmt.Println(b)
	fmt.Println(global)
}
