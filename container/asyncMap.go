package main

import (
	"fmt"
	"sync"
	"time"
)

//下面来看一个示例引入要介绍的内容

func fail() {
	var m6 = make(map[string]string)

	//两个协程不停的读写map,运行代码会报错concurrent map read and map write
	go func() {
		for {
			m6["1"] = "1"
		}
	}()
	go func() {
		for {
			_ = m6["1"]
		}
	}()
	time.Sleep(10 * time.Second)
}

func success() {
	//如果需要在不同的协程对同一个map对象进行操作,可以试试sync.Map
	var m7 sync.Map
	m7.Store(1, 2)
	m7.Store("3", 4)        //赋值操作
	fmt.Println(m7.Load(3)) //取值操作
	m7.Delete(1)            //删除键值对
	m7.LoadOrStore("3", 5)  //不存在就创建,存在就不做任何操作
	fmt.Println(m7.Load("3"))

	//遍历它,好奇怪的操作
	m7.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	//使用sync.Map不会出现异常
	var m6 sync.Map
	go func() {
		for {
			m6.Store("1", 1)
		}
	}()
	go func() {
		for {
			m6.Load("1")
		}
	}()
	time.Sleep(10 * time.Second)

}

func main() {
	//fail()
	success()

}
