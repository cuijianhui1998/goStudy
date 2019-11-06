package main

import "fmt"

func main() {
	//指针表示只想变量的地址,有两个操作符,&:取变量的地址;*:指针
	var a = 5
	var ptr *int = &a //ptr的值就是a的地址,ptr是一个int型的指针变量
	fmt.Println(ptr)
	var c int = *ptr //取这个地址所指向的值,c为5
	fmt.Println(c)

	//创建指针的另一个方法new
	str := new(int) //默认值为nil值
	fmt.Println(*str)

}
