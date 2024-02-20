package main

import "fmt"

func main() {
	str1 := [][]string{
		[]string{"C Prog", "FYBCA"},
		[]string{"CPP Prog", "SYBCA"},
		[]string{"GO Prog", "TYBCA"},
	}
	fmt.Println("MultiDimensional Slice str1:")
	for i := 0; i < len(str1); i++ {
		fmt.Println(str1[i])
	}
	fmt.Println("Slice str1 like Matrix:")
	n := len(str1)
	m := len(str1[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(str1[i][j] + " ")
		}
		fmt.Print("\n")
	}
}
