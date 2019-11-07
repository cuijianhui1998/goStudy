package main

import "fmt"

//映射类型

func main() {
	var m1 = map[string]string{"hello": "good"}
	var m2 map[string]string
	//m2["good"] = "hello" 无法向空map中添加k-v键值对
	m1["good"] = "hello"
	fmt.Println(m1, m2)

	m3 := make(map[string]string, 4)
	m3["good"] = "hello" //通过这种方式可以向一个没有任何k-v键值对的map中添加值
	fmt.Println(m3)

	//map的遍历,类似于python中dict的items()
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//删除一个键值对,没有清空的函数,垃圾回收比清空一个map要高效
	delete(m1, "hello")

}
