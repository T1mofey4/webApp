package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"username"`
}

func main() {
	// user1 := User{Name: "Ivan", Id: 222}
	// bytes, _ := json.Marshal(user1)
	// bytes1, _ := json.MarshalIndent(user1, "", "  ") //second way with Tab
	// fmt.Println(string(bytes))
	// fmt.Println(string(bytes1))

	http.HandleFunc("/user", UserHandler)
	http.ListenAndServe(":3000", nil)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	user1 := User{Name: "Ivan", Id: 222}
	bytes, err := json.Marshal(user1)
	if err != nil {
		w.Write([]byte("500 Internal Server Error" + err.Error()))
		return
	}
	w.Header().Set("ContentType", "application/json")
	w.Write(bytes)
}
