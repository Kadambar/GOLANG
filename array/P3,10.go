package main

import "fmt"

func main() {
	a := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println("Number of rows in array: %d\n", len(a))
	fmt.Println("Number of columns in array: %d\n", len(a[0]))
	fmt.Println("Total number of elements in array: %d\n", len(a)*len(a[0]))
	fmt.Println("Traversing Array")
	for _, row := range a {
		for _, val := range row {
			fmt.Println(val)
		}
	}
}
