package main

import (
	"fmt"
	"math/rand"
	"time"
)

func product(header string, channel chan<- string)  {
	//chan<- 用于向通道发送数据
	for{
		channel<-fmt.Sprintf("%s:%v", header,rand.Int31())
		time.Sleep(time.Second)
	}
}

func custom(channel <-chan  string)  {
	// <-chan 用于接收数据
	for  {
		res := <- channel
		fmt.Println(res)

	}
}

func main(){
	channel := make(chan string)
	go product("dog", channel)
	go product("cat", channel)
	custom(channel)
}
