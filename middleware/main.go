package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/foo", FooHandler)

	handler := MyMiddleWare(mux)
	handler = SecondMiddleWare(handler)

	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		panic(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
	w.Write([]byte("Home"))
}

func FooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("foo")
	w.Write([]byte("foo"))
}

func SecondMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before2")
		handler.ServeHTTP(w, r)
		fmt.Println("after2")
	})
}

func MyMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before")
		handler.ServeHTTP(w, r)
		fmt.Println("after")
	})
}
