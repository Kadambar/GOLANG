package main

import "fmt"

func main() {
	str := []string{
		"Radha",
		"Pallawi",
		"Minal",
		"Kusum",
	}
	S1 := make([]string, len(str))
	copy(S1, str)
	fmt.Println("Original Array")
	fmt.Println(str)
	fmt.Println("Slice  S1")
	fmt.Println(S1)
}
