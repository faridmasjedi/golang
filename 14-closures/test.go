package main

import "fmt"

func intSeq(a int) func() int {
	return func() int {
		a++
		return a
	}
}

func sumFromOneToANum(x int) func() int {
	return func() int {
		return x * (x + 1) / 2
	}
}

func multiple(x int) func(y int) int {
	return func(y int) int {
		return x*y
	}
}

func main() {
	nextInt := intSeq(3)
	println(nextInt())

	sumFrom1To10 := sumFromOneToANum(10)
	fmt.Println(sumFrom1To10())
	multipleWith3 := multiple(3)
	
	for i:=1; i<10; i++ {
		fmt.Println("Number:", i , "|", "multiply:", multipleWith3(i))
	}
}
