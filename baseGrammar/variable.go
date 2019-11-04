package main

import "fmt"

//变量的声明与初始化

//定义常量
const Pi = 3.13

func add(c int, g int) int {
	return c + g
}

func main() {
	//变量的定义有两种
	//1. 使用var关键字
	var a int
	var b int         //未初始化的默认为nil值
	var c, d = 1, "a" //如果直接初始化可以不显示指明类型,会根据等号右边自动判断
	var (
		//中间没有逗号
		e int
		f string
	)
	fmt.Print(a, b, c, d, e, f)
	//2. 使用:=符号定义,会根据右侧的值自动判断变量的类型,但是只能在函数内使用
	g := 5
	fmt.Println(e)

	//匿名变量,可以参考python中的占位符,add有一个返回值,但是我们可以使用_不接受占位不接收
	//好处:对于一些我们不需要的返回值,我们可以抛弃它,不占内存空间
	_ = add(c, g)
}
