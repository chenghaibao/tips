package main

import "net/http"

func main() {

	// 偷懒方式
	//http.ListenAndServe("127.0.0.1:8080", nil)
	//http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("hello world!"))
	//})


	// 指针方式
	server := &http.Server{
		Addr: "127.0.0.1:8080",
	}

	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})

	server.ListenAndServe()

}