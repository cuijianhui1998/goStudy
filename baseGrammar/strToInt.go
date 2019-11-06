package main

import (
	"fmt"
	"strconv"
)

func main() {
	//虽然显示转换中,字符串无法转换为int型,但是提供了一个内置模块strconv
	var a = 8
	var b = "5"
	//字符串转int,第二个返回值是error类型,由于str转int有可能会失败,所以才会有第二个返回值
	var c, _ = strconv.Atoi(b)
	fmt.Print(c)
	//int转字符串,很奇怪的是这里返回的确实一个返回值,疑惑(疑惑已解决,str转int可能失败)
	var d = strconv.Itoa(a)
	fmt.Print(d)

	//str和其他类型的转换有Parse(),Format()
	//Append()用于将指定类型转为字符串后加入到切片中

}
