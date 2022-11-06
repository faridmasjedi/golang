package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func main(){
	r := rect{width: 3, height: 4}
	fmt.Println("area: ", r.area())
}