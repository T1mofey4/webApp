package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"username"`
}

func main() {
	user1 := User{Name: "Ivan", Id: 222}
	bytes, _ := json.Marshal(user1)
	bytes1, _ := json.MarshalIndent(user1, "", "  ") //second way with Tab
	fmt.Println(string(bytes))
	fmt.Println(string(bytes1))
}
