package main

import (
	"errors"
	"fmt"
)

type errorer interface {
	Error() string
}

type MyError struct{}

func (e MyError) Error() string {
	return fmt.Sprintf("自定义错误类")
}

func main() {
	err := errors.New("错误")
	fmt.Printf("%T, %s\n", err, err)
	var err1 = fmt.Errorf("错误1")
	fmt.Printf("%T, %s\n", err1, err1)
	var err2 error
	err2 = MyError{}
	fmt.Printf("%T, %s\n", err2, err2)

	if err, ok := err2.(errorer); ok {
		fmt.Println(err, ok)
	}

}
