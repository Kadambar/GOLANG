package main

import "fmt"

func main() {
	total := 0
	ids := make([]int, 5)
	/*for_index, value:=range ids{
	total = total+value*/
	for _, value := range ids {
		total = total + value
	}
	fmt.Println("Total of all ids = ", total)
}
