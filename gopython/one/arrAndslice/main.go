package main

import "fmt"

func main() {
	str1 := "hello世界!"
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len([]rune(str1)), len(arr1))

	// 切片定义的三种方法
	arr2 := [1]int{1}
	slice1 := arr2[:]
	var slice2 []int
	slice3 := make([]int, 0, 10)
	slice3 = append(slice3, 1)
	fmt.Println(slice1, slice2, slice3)
	fmt.Printf("%T, %v, %v\n", slice1, slice1, cap(slice1))

	// 当数组切片扩容之后,指向的内存地址会发生变化, 底层就不在是引用的数组
	slice1[0] = 8
	fmt.Println(arr2, slice1)
	slice1 = append(slice1, 2)
	fmt.Println(arr2, slice1)
	slice1[0] = 9
	fmt.Println(arr2, slice1)

}
