package main

import "fmt"

func main() {
	a := [2][3]int{{10, 20, 30}, {40, 50, 60}}
	fmt.Println("Using for-range")
	for _, row := range a {
		for _, val := range row {
			fmt.Println(val)
		}
	}
	fmt.Println("\nUsing for loop")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(a[i][j])
		}
	}
	fmt.Println("\nUsing for loop - Second way")
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Println(a[i][j])
		}
	}
}
