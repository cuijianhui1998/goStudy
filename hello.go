package main

import (
	"fmt"
	"math/rand"
	"time"
)

//为并发而生的golang

func product(header string, channel chan<- string) {
	//chan<- 用于向通道发送数据
	for {
		channel <- fmt.Sprintf("%s:%v", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

func custom(channel <-chan string) {
	// <-chan 用于接收数据
	for {
		res := <-channel
		fmt.Println(res)

	}
}

func main() {
	//一个典型的生产者-消费者模型,通过协程的模式,使用通道进行数据通信
	channel := make(chan string)
	go product("dog", channel)
	go product("cat", channel)
	custom(channel)
}
