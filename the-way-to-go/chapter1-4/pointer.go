package main

import (
	"fmt"
)

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)

	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	// intP 存储了 i1 的内存地址；它指向了 i1 的位置，它引用了变量 i1
}
