package main

import "fmt"

func main() {
	for {
		var score int
		fmt.Println("请输入分数")
		fmt.Scanln(&score)
		if score == -1 {
			break
		}

		switch {
		case score >= 90:
			fmt.Println("A+")
		case score >= 80:
			fmt.Println("A")
		case score >= 70:
			fmt.Println("B+")
		case score >= 60:
			fmt.Println("B")
		default:
			fmt.Println("C")
		}
	}

}
