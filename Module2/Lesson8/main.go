package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Hello")
}

func main() {
	printHello()
	fmt.Println("GoodBye")
	time.Sleep(time.Second)
}
