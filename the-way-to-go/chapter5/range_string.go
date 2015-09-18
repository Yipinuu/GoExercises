package main

import (
	"fmt"
)

func main() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("%d : %c \n", pos, char)
	}
	fmt.Println()
	str2 := "Chinese: 瓷器"
	fmt.Printf("The length of str2 is %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("%c : %d \n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune char bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d %d %U '%c' %X\n ", index, rune, rune, rune, []byte(string(rune)))
	}
}
