package main

import "fmt"

func fun(a [3]int, size int) int {
	var i, val, r int
	for i = 0; i < size; i++ {
		val += a[i]
	}
	r = val / size
	return r
}
func main() {
	var arr = [3]int{1, 9, 2}
	var res int
	res = fun(arr, 3)
	fmt.Printf("Final result is: %d ", res)
}
