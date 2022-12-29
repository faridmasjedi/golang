package main

import "fmt"

func main(){
	if 4 % 2 == 0 {
		fmt.Println("4 is even")
	}else{
		fmt.Println("4 is odd")
	}


	fmt.Println("\n---\n")

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	}else if num > 9 {
		fmt.Println(num, "is positive & greater than 9")
	} else {
		fmt.Println(num, "is positive & less than 9 or equal to 9")
	}
}