package main

import "net/http"

func main() {
	http.ListenAndServe(":500", http.FileServer(http.Dir("wwwroot")))
}
