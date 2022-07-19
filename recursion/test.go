package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibo(n int) int {
	if n < 2 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {
	fmt.Println(factorial(3))
	fmt.Println(factorial(5))
	fmt.Println(fibo(5))
}
