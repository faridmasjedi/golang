package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	fmt.Print("Write ", i , " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println("\n----\n")

	var day = time.Now().Weekday()
	switch day {
	case time.Saturday, time.Sunday: 
		fmt.Println("It is weekend")
	default:
		fmt.Println("It is weekday")
	}

	fmt.Println("\n----\n")

	var t = time.Now()
	switch {
	case t.Hour() < 12 :
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	}
}