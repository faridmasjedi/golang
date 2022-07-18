package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, v := range nums {
		total += v
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 3, 5)

	num := []int{2, 3, 4, 6}
	sum(num...)
}
