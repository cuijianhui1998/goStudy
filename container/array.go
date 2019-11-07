package main

import "fmt"

//数组的特性:定长,而且还必须是只能包含一种数据类型

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3

	var b = [3]int{1, 2, 3}
	fmt.Print(b[0])

	//遍历数组,index代表的是索引,i代表的是数组对应索引的值
	//在range中为每一个元素提供了一个副本,所以i和b[index]其实并不是同一个地址
	for index, i := range b {
		fmt.Print(index, i)
	}

	//声明多维数组
	var c = [4][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	fmt.Println(c[3][1])

}
