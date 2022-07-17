package main

import "fmt"

func main(){
	var a [5]int
	fmt.Println("a: ", a , "\n----\n")

	a[4] = 100
	fmt.Println("new a: ", a , "\n----\n")

	fmt.Println("len of a: ", len(a), "\n---\n")

	b := [5]int{2,4,6,8,10}
	fmt.Println("b is: ", b, "\n----\n")

	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 2; j < 5; j++ {
			c[i][j-2] = i + j
		}
	}
	fmt.Println("c is : ", c)
 }