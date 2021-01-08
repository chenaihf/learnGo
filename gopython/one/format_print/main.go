package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("请输入你的姓名及年龄")
	fmt.Scanln(&name, &age)
	summary := fmt.Sprintf("我叫%s, 今年%d岁了", name, age)
	fmt.Println(summary)
	fmt.Printf("%+v type:%T, unicode:%c\n", name, name, name)
	fmt.Printf("%+v type:%T, unicode:%c\n", age, age, age)
}
