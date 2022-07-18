package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("map: ", m)
	fmt.Println("map len: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	_, pr := m["k1"]
	fmt.Println("pr: ", pr)

	n := map[string]int{"s1": 1, "s2": 2, "s3": 3}
	fmt.Println("map n: ", n)
}
