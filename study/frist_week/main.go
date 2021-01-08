package main

import "fmt"

func main()  {
	fmt.Println("hello world!")

	intvar := 10
	floatvar := 3.14
	//bvar := byte('9')
	// 单引号是byte类型, 双引号是字符串类型
	var btvar = byte(1)
	bvar := '你'
	name := "chenai"
	fmt.Println(intvar, floatvar, bvar, btvar, name)
	fmt.Printf("%T, %T, %T, %T, %T \n", intvar,floatvar,bvar,btvar,name)

	//数组
	fmt.Println("数组:")
	var arr [10]int
	arr[1] = 10
	fmt.Println(arr)
	arr1 := [10]int{1,2,3,4,5}
	fmt.Println(arr1)

	//切片
	//切片不会被定义时的长度限制
	fmt.Println("切片:")
	slice1 := make([]int, 1)
	fmt.Println(slice1)
	slice1 = arr1[:5]
	fmt.Println(slice1)
	slice2 := make([]interface{}, 10)
	slice2[1] = "你好"
	slice2[3] = 1
	slice2[5] = 0.2
	slice2[9] = true
	fmt.Println(slice2)
	slice3 := []int{1,2,3}
	fmt.Println(slice3)

	//map
	fmt.Println("map 字典")
	map1 := make(map[string]string)
	map1["小明"] = "xiaoming"
	map1["小红"] = "xiaohong"
	fmt.Println(map1["小明"])
	clienttype := make(map[int]string)
	clienttype = map[int]string{100:"user", 101:"phone", 200:"weixin"}
	fmt.Println(clienttype)
	for typ := range clienttype{
		fmt.Println(typ, clienttype[typ])
	}
	//
	val, ok := clienttype[100]
	if ok{
		fmt.Println(val, ok)
	}else {
		fmt.Println(ok)
	}
	//go特有
	if val, ok := clienttype[300];ok{
		fmt.Println(val)
	}else {
		fmt.Println("not found")
	}
	//map切片, map是引用数据类型, 定义的数据都是指针,所以不需要用切片来指向它
	//mapslice := make([]map[int]string, 2)
	//mapslice = clienttype[:2]






}
