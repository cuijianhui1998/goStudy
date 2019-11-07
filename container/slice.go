package main

import "fmt"

//切片类型是一种很有意思的类型

func main() {
	//与数组的初始化只是少了一个长度
	var a = []int{1, 2, 3}
	fmt.Print(a, len(a))

	//对数组进行切片生成一个slice对象
	var b = [5]int{1, 2, 3, 4, 5}
	var c = b[:4]
	fmt.Print(c)

	//使用make函数生成一个切片, 第二个参数表示给切片初始化前n个值,默认值nil
	//第三个参数是预分配的容量
	var d = make([]int, 2, 10)
	fmt.Print(d)

	//append向切片中添加元素,需要注意的是python中的append不会重新生成一个对象,
	// 但是在golang中却不是原地操作
	var e = []int{1, 2, 3}
	f := append(e, 4)
	g := append([]int{0}, a...) //向开头添加元素(不得不吐槽一下,不仅丑陋,而且与前面的接口完全不同)
	fmt.Println(f, g)
	//也有类似于python中extends操作,同样,他也不是原地操作
	//...表示对切片解包,可以类比python中元组的解包
	var h = []int{4, 5, 6}
	var i = append([]int{}, h...)
	fmt.Println(i)
	//终于明白上面向切片的前面添加元素如此的丑陋,
	// 原理其实是在容量为一个元素的切片中append原切片

	//切片的复制,将k的值复制给j,长度只会取较短切片的值,
	// 如果源切片长度不够,会保留切片原有位置的值(第一个参数是目标切片,第二个参数是源切片)
	var j = []int{1, 2, 3}
	var k = []int{4, 5, 6, 7, 8, 9}
	copy(j, k)
	//copy(k, j)
	fmt.Println(j, k)

	//删除操作,可以类比python中对list的切片操作都不是原地操作

}
