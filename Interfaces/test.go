package main

import (
	"fmt" 
	"math"
)

type rect struct {
	width, height float64
}

type circle struct {
	raduis float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) prim() float64{
	return 2*r.width*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.raduis * c.raduis
}
func (c circle) prim() float64 {
	return 2 * math.Pi * c.raduis
}

type geometry interface {
	area() float64
	prim() float64
}

func measure(g geometry){
	fmt.Println("g:",g)
	fmt.Println("area:",g.area())
	fmt.Println("prim:",g.prim())
	fmt.Println("\n")
}

func main(){
	r := rect{width:3, height:4}
	c := circle{raduis: 1}

	measure(r)
	measure(c)
}