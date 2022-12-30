package main

func vals() (int, int) {
	return 3, 7
}

func mathOp(a int, b int) (int, int) {
	return a + b, a * b
}

func main() {
	a, b := vals()
	x, y := mathOp(a, b)
	println("a:", a, " | b:", b)
	println("a+b:", x, " | a*b:", y)
}
