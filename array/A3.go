package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a [3]int
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

}
