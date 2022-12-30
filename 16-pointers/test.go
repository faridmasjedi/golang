package main

import "fmt"

var i int = 0

func withoutPointer(ival int) {
	ival = 1
}

func withPointer(iptr *int) {
	*iptr = 1
}

func main() {
	fmt.Println("ival:", i)

	withoutPointer(i)
	fmt.Println("without pointer | i:", i)

	withPointer(&i)
	fmt.Println("with pointer | i:", i)
}
