package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	//创建web服务
	http.ListenAndServe("localhost:1000", nil)
}
