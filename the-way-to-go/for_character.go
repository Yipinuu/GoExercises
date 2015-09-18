package main

import (
	"fmt"
)

func main() {
	// 1
	for i := 0; i < 25; i++ {
		for j := 0; j < i; j++ {
			print("G")
		}
		fmt.Println()
	}

	// 2
	for str, i := "G", 0; i < 25; i++ {
		fmt.Printf("%s\n", str)
		str += "G"
	}
}
