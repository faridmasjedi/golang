package main

import (
	"fmt"
	"math"
)

const s = "String s"
const ss string = "String ss"

func main(){
	fmt.Println( "s value is: " + s + " , and ss values is: " + ss)
	const n = 500

	const d = 3e3
	const ff = d/n
	var sins = math.Sin(n)

	fmt.Println(n)
	fmt.Println(d)
	fmt.Println(ff)
	fmt.Println(int64(ff))
	fmt.Println(sins)
}