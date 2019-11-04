package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("."))) //http服务器
	http.HandleFunc("/index", indexHandler)          //视图函数和路径相对应
	http.ListenAndServe(":8000", nil)
}
