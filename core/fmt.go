package core

import "fmt"

func Println(messages ...interface{}) {
	fmt.Println(messages...)
}

func Printf(params ...interface{}) {
	fmt.Printf(fmt.Sprintf("%v", params[0]), params[1:]...)
}
