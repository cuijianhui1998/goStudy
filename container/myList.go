package main

import (
	"container/list"
	"fmt"
)

func main() {

	var l list.List //这是一个双链表
	var l1 list.List
	l1.PushBack(3)
	l.PushBack(1)
	l.PushBack("2") //向后插入元素,python中list的append
	l.PushFront(0)  //向前插入元素
	l.PushBackList(&l1)
	a := l.Back() //是一个element对象,在删除时用得上
	fmt.Println(*a)
	fmt.Println(l)
	l.Remove(a)
	//遍历双链表,从头结点开始遍历
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	//从尾节点开始遍历
	for i := l.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}
}
