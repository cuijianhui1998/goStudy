package main

import "fmt"

func main() {
	//不存在隐式类型转换,所以必须显示的转换类型
	//小精度转大精度没问题,大精度转小精度可能会出错
	//下方的b值最后是-128
	var a int16
	a = 128
	var b = int8(a)
	fmt.Print(b)

	//底层不同类型之间无法相互转换,例如整型和字符串之间无法强转

}
