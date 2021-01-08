package main

import "fmt"

func div(a, b int) int {
	if b == 0 {
		panic("除数不能为0")
	}

	// recover是捕获不到子线程抛出的错误的
	go func() {
		//defer func() {
		//	err := recover()
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//}()
		panic("子线程错误")
	}()

	return a / b
}

func main() {
	// 捕获异常的写法
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	res := div(1, 1)

	fmt.Println(res)
	fmt.Println("break")

}
