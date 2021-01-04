package main

import "net/http"

// 自定义Handler
type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

type helloGo struct{}

func (m *helloGo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello Golang"))
}
func main() {
	m := helloHandler{}
	g := helloGo{}
	server := http.Server{
		Addr:    "localhost:2000",
		Handler: nil, // 如果是nil，会启动多路由
	}

	// 自定义路由
	http.Handle("/hello", &m)
	http.Handle("/go", &g)
	server.ListenAndServe()
	// 两种方法作用相同，上方比较灵活可以自由配置
	//http.ListenAndServe("localhost:2000",nil)
}
