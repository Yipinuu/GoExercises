package main

import (
  "fmt"
  "unicode/utf8"
)

func main() {
  // count number of characters:
  str1 := "asSASA ddd aksjdlldwajdk saL"
  fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
  fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
  str2 := "asSASA ddd aksjdlldwajdk你猜 saL"
  fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
  fmt.Printf("The number of characters in string str2 is %d\n", utf8.RuneCountInString(str2))
}