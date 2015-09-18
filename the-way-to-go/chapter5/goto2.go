// compile error goto2.go:10: goto TARGET jumps over declaration of b at ./goto2.go:11
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
}
