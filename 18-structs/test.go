package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) person {
	var p person
	p.name = name
	p.age = age
	return p
}

func main() {
	x := person{name:"Farid", age:55}
	fmt.Println("x:",x)

	s := newPerson("Farid", 37)
	fmt.Println("s:", s)

	ss := s
	ss.age = 65

	fmt.Println("ss:", ss)
	fmt.Println("s:", s)

	sss := &ss
	sss.age = 56
	fmt.Println("------")
	fmt.Println("ss:", ss)
	fmt.Println("sss:", sss)

}
