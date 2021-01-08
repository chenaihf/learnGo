package main

import (
	"ccGo/retriever/real"
	"fmt"
)

type Retriever interface {
	Get(string) string
}

func download(r Retriever) string{
	return r.Get("https://www.imooc.com")
}

func main() {
	r := real.Retriever{}
	//res := r.Get("https://www.imooc.com")
	res := download(r)
	fmt.Println(res)
}
