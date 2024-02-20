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
	S1 := append(str, "Satish")
	fmt.Println(S1)
	fmt.Println(str)
}
