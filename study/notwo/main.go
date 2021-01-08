package main

import "fmt"

func main() {
	fmt.Println(mysum(4))
	fmt.Println(fbnc(11))
}

func mysum (num int) int{
	if num == 1{
		return 1
	}else {
		return mysum(num-1) + num
	}
}

func fbnc (idx int) int{
	if idx == 1 || idx ==2{
		return 1
	}else {
		return fbnc(idx-1) + fbnc(idx-2)
	}
}
