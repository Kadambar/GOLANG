package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	b := a
	println("The Original array a is...")
	fmt.Println(a)
	println("The copied array b is...")
	fmt.Println(b)
	b[0] = 100
	println("The Original array a is...")
	fmt.Println(a)
	println("The copied array b is...")
	fmt.Println(b)
}
