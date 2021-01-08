package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
	fmt.Printf("%p: %v, %p: %v\n", &a, a, &b, b)
}

func swapPointer(a, b *int) {
	*a, *b = *b, *a
	// 指针变量也会开辟一块内存存放传入的指针地址, 所以这里打印的内存地址不是传入变量的内存地址
	// 指针变量只是储存某一个值的内存地址, 跟值本身没有关联, 但是可以通过 *运算符取得变量中内存地址对应的值
	fmt.Printf("%p: %v, %p: %v\n", &a, a, &b, b)
}

func main() {
	a := 1
	b := 9
	swap(a, b)
	fmt.Printf("%p: %v, %p: %v\n", &a, a, &b, b)
	swapPointer(&a, &b)
	fmt.Printf("%p: %v, %p: %v\n", &a, a, &b, b)
	fmt.Println("-----------------------------------")

	var p1 *int
	fmt.Printf("%T, %v, %p\n", p1, p1, &p1)
	//*p1 = 10
	p1 = &a
	fmt.Println(p1, "   ", *p1)
	*p1 = 10
	fmt.Println(p1, "   ", *p1)
	fmt.Println("-----------------------------------")

	// new函数, 给指针变量一个初始值(一个内存地址), 有了初始的地址(指针变量存放的是内存地址), 所以肚子里要有内存地址才可以使用*运算符
	// new函数放回一个类型初始值的地址
	var p2 *int = new(int)
	fmt.Printf("%T, %v, %p\n", p2, *p2, &p2)
	*p2 = 10
	fmt.Println(p2, "   ", *p2)
	fmt.Println("-----------------------------------")

	// make函数, map类型必须使用make声明或者声明时直接赋值
	var info map[string]string
	fmt.Printf("%T, %v, %p\n", info, info, &info)
	infoM := make(map[string]string)
	fmt.Printf("%T, %v, %p\n", infoM, infoM, &infoM)
	//info["A"] = "1"
	infoM["A"] = "1"
	fmt.Printf("%T, %v, %p\n", info, info, &info)
	fmt.Printf("%T, %v, %p\n", infoM, infoM, &infoM)
	fmt.Println("-----------------------------------")

}
