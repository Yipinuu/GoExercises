package main

import (
  "fmt"
  "strings"
)

func main() {
  var str string = "Hi, I'm Marc, Hi."

  fmt.Printf("This position of \"Marc\" is: ")
  fmt.Printf("%d\n", strings.Index(str, "Marc"))

  fmt.Printf("This position of first instance of \"Hi\" ")
  fmt.Printf("%d\n", strings.Index(str, "Hi"))
  fmt.Printf("This position of last instance of \"Hi\" ")
  fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

  fmt.Printf("The position of \"Burger\" is: ")
  fmt.Printf("%d\n", strings.Index(str, "Burger"))
}
