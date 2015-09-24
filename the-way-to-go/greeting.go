package main

func main() {
	println("In main before call greeting")
	greeting()
	println("in main after all call greeting")
}

func greeting() {
	println("In greeting: Hi!!!!")
}
