package main

import "fmt"

func main() {
	str := []string{
		"Radha",
		"Pallawi",
		"Minal",
		"Kusum",
	}
	fmt.Println("Before Append")
	fmt.Println(str)
	str = append(str, "Satish", "Niket", "Dipali")
	fmt.Println("After Append")
	fmt.Println(str)
}
