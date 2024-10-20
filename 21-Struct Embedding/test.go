package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Println("num:", co.num)
	fmt.Println("num:", co.base.num)
	fmt.Println("describe co:", co.describe())
	fmt.Println("describe co:", co.base.describe())
}
