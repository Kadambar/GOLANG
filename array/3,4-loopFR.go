package main

import "fmt"

func main() {
	count := 1
	var a = make([][]int, 4)
	for i := 0; i < 4; i++ {
		a[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			a[i][j] = count
			count++
		}
	}
	for index, value := range a {
		fmt.Println(index, "=", value)
	}
}
