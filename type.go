package main
import "fmt"

type TZ int

func main() {
  var a,b TZ = 3,4
  var c := a + b
  fmt.Printf("c has the value: %d", c)// c has the value: 7
}
