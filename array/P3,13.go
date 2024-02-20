package main

import "fmt"

func main() {
	a := [2][3]byte{}
	fmt.Println("first row")
	fmt.Println(&a[0][0])
	fmt.Println(&a[0][1])
	fmt.Println(&a[0][2])
	fmt.Println("second row")
	fmt.Println(&a[1][0])
	fmt.Println(&a[1][1])
	fmt.Println(&a[1][2])
}
