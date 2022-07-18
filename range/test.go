package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	var sum int = 0

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index ", i, " : ", num)
		}
		sum += num
	}

	println("sum: ", sum)

	for k, v := range map[string]int{"s1": 1, "s2": 2, "s3": 3} {
		fmt.Println("key: ", k, " | value: ", v)
	}

	for k := range map[string]int{"ss1": 1, "ss2": 2} {
		fmt.Println("just key: ", k)
	}

	for i, c := range "Masjedi" {
		println("key: ", i, " | value: ", c)
	}

}
