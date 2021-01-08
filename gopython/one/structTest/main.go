package main

import (
	"encoding/json"
	"fmt"
)

type people struct {
	Name   string `json:"name"`
	Age    int    `json:"age, omitempty"`
	Gender string `json:"gender"`
}

func main() {
	my := people{
		Name:   "尘埃",
		Age:    18,
		Gender: "男",
	}
	myjson, _ := json.Marshal(my)
	fmt.Printf("%T, %T\n", myjson, myjson[1])
	fmt.Println(string(myjson))
}
