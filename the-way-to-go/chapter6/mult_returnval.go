package main

var n1, n2 int

func main() {
	n1, n2 = func1(2, 3)
	println(n1, "\n", n2)
	n1, n2 = func2(2, 3)
	println(n1, "\n", n2)
}

func func1(input int, input2 int) (int, int) {
	return input + input2, input * input2
}

func func2(input int, input2 int) (i int, p int) {
	i = input + input2
	p = input * input2
	return i, p
}

// package main

// import (
//     "fmt"
// )

// func SumProductDiff(i, j int) (int, int, int) {
//     return i+j, i*j, i-j
// }

// func SumProductDiffN(i, j int) (s int, p int, d int) {
//   s, p, d = i+j, i*j, i-j
//     return
// }

// func main() {
//     sum, prod, diff := SumProductDiff(3,4)
//     fmt.Println("Sum:", sum, "| Product:",prod, "| Diff:", diff)
//   sum, prod, diff = SumProductDiffN(3,4)
//     fmt.Println("Sum:", sum, "| Product:",prod, "| Diff:", diff)
// }
