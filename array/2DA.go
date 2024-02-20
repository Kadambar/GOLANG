package main

import "fmt"

func main() {
	a := [2][3]int{{10, 20, -3}, {4, 8, 16}}
	fmt.Println("Rows in array a are: ", len(a))
	fmt.Println("Columns in array a are: ", len(a[0]))
	fmt.Println("Total elements in array a are: ", len(a)*len(a[0]))
	for _, row := range a {
		for _, val := range row {
			fmt.Print(val, " ")
		}
	}
}
