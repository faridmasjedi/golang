package main

import (
	"fmt"
	"time"
)

func sync(msg chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("done!!!")

	msg <- true
}

func main() {
	d := make(chan bool, 1)

	go sync(d)
	<-d
}
