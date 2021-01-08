package main

import (
	"fmt"
	"strings"
)

func main() {
	defer fmt.Println("100")
	p1 := person{"张三", 18, address{"重庆"}}
	p1.print()
	fmt.Println(p1.city)
	fmt.Println(p1.name)
	f1 := fbnc()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	f2 := sum()
	fmt.Println(f2(10))
	fmt.Println(f2(10))
	f3 := validata(".jpg")
	fmt.Println(f3("hello"))
	fmt.Println(f3("hellofasd.png"))
}

type person struct {
	name string
	age uint
	address
}

type address struct {
	city string
}

func (addr address) print() {
	fmt.Println("hello address")
}

func fbnc() func()int{
	a,b := 0,1

	return func() int {
		a, b = b, a+b
		return a
	}
}

func sum() func(int)int{
	sum := 0

	return func(num int) int {
		sum = num+sum
		return sum
	}
}

func validata(vali string) func(string)string{

	return func(s string) string {
		if strings.HasSuffix(s, vali){
			return s
		}else {
			return s + vali
		}
	}
}
