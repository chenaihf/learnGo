package main

import (
	"fmt"
	"strconv"
)

func main() {
	const PI = 3.14
	bstr := strconv.FormatBool(false)
	itoa := strconv.Itoa(123)
	atoi, err := strconv.Atoi("999")
	if err != nil {
		fmt.Println(err)
	}
	fstr := strconv.FormatFloat(PI, 'e', -1, 64)
	fmt.Printf("%T  %v \n", bstr, bstr)
	fmt.Printf("%T  %v \n", itoa, itoa)
	fmt.Printf("%T  %v \n", atoi, atoi)
	fmt.Printf("%T  %s \n", fstr, fstr)
}
