package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/foo", FooHandler)

	r.Use(MyMiddleWare)
	r.Use(SecondMiddleWare)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home gorilla")
	w.Write([]byte("Home"))
}

func FooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("foo gorilla")
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
