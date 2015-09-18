package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var orig string = "ABC"
	var nesW string

	fmt.Printf("The size of ints is : %d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an ntger - exiting with error\n", orig)
		os.Exit(1)
	}
	fmt.Printf("The integer is %d\n", an)
}
