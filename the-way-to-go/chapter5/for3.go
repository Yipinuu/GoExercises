package main

import (
	"fmt"
)

func main() {
	for {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
		if i < 0 {
			breck
		}
	}
}
