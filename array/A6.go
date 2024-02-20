package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [...]int{1, 2, 3}
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
	var i int
	for i < len(a) {
		fmt.Print(a[i], " ")
		i++
	}
	fmt.Println()
	for _, val := range a {
		fmt.Print(val, " ")
	}
	fmt.Println(reflect.ValueOf(a))
}
