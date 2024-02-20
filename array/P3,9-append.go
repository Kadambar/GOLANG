package main

import "fmt"

func main() {
	str := []string{
		"Radha",
		"Pallawi",
		"Minal",
		"Kusum",
	}
	fmt.Println(str)
	str = append(str, "Satish")
	fmt.Println(str)
}
