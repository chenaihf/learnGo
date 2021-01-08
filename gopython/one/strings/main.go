package main

import (
	"fmt"
	"strings"
)

func main() {
	//1. 是否包含某个子串
	var name string = "bobby:\"慕课网\""
	fmt.Println(strings.Contains(name, "慕课网"))
	fmt.Println(strings.Index(name, "慕课网"))

	//2. 统计出现的次数
	fmt.Println(strings.Count(name, "b"))

	//3. 前缀和后缀
	fmt.Println(strings.HasPrefix(name, "o"))
	fmt.Println(strings.HasSuffix(name, "\""))

	//4. 大小写转换
	fmt.Println(strings.ToUpper("bobby"))
	fmt.Println(strings.ToLower("BOBBY"))

	//5. 字符串的比较
	fmt.Println(strings.Compare("ab", "aa")) //字符的比较就是ascii的比较 返回-1， 1， 0
	fmt.Println(strings.Compare("b", "a"))   //字符的比较就是ascii的比较 返回-1， 1， 0
	fmt.Println(strings.Compare("b", "b"))   //字符的比较就是ascii的比较 返回-1， 1， 0

	//6. 去掉空格和指定字符串
	fmt.Println(strings.TrimSpace(" bobby "))
	fmt.Println(strings.TrimLeft("bobby", "b"))
	fmt.Println(strings.Trim("bobby", "b"))

	//7. split方法
	fmt.Println(strings.Split("bobby imooc", " "))

	//8. 合并 join方法将字符串数组连接起来
	arrs := strings.Split("bobby imooc", " ")
	fmt.Println(strings.Join(arrs, "-"))

	//9. 字符串替换
	fmt.Println(strings.Replace("bobby: 18 电话：18888888", "18", "19", 2))

	fmt.Printf("%T, %T, %v, %c\n", 'a', name[1], name[1], name[1])
	fmt.Printf("%T, %T, %v, %c\n", '1', "中", name[1], name[1])

}
