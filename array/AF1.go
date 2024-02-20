package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	b := a[:2]
	c := a[1:3]
	d := a[2:]
	e := a[:]
	fmt.Println(a, reflect.TypeOf(a))
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(c, reflect.TypeOf(c))
	fmt.Println(d, reflect.TypeOf(d))
	fmt.Println(e, reflect.TypeOf(e))

}
