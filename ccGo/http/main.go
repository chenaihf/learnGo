package main

import (
	"fmt"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello http!")
}

func main() {
	http.HandleFunc("/",sayhello )
	http.ListenAndServe(":8080", nil)
	//if nil
}
