package main

import (
	"fmt"
	"sort"
)

type Course struct {
	name  string
	price int
}

type Courses []Course

func (c Courses) Len() int {
	return len(c)
}

func (c Courses) Less(i, j int) bool {
	return c[i].price < c[j].price
}

func (c Courses) Swap(i, j int) {
	c[i].price, c[j].price = c[j].price, c[i].price
}

func main() {
	c := Course{"go", 100}
	courses := Courses{
		Course{"python", 150},
		Course{"java", 200},
		c,
	}
	fmt.Println(courses)
	sort.Sort(courses)
	fmt.Println(courses)
}
