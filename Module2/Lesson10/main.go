package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func gor1(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "gor1"
}

func gor2(ch chan string) {
	time.Sleep(time.Second * 3)
	ch <- "gor2"
}

func main() {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-sigChan
		fmt.Printf("Gracefull shutdown...%v", sig)
		time.Sleep(time.Second * 2)
		os.Exit(0)
	}()

	fmt.Println("Programm is working press hotkeys")
	select {}

	// ch1 := make(chan int)

	// ch1 <- 10

	// fmt.Println(<-ch1)

	// 	ch1 := make(chan string)
	// 	ch2 := make(chan string)

	// 	go gor1(ch1)
	// 	go gor2(ch2)

	// 	select {
	// 	case res := <-ch1:
	// 		fmt.Println(res)
	// 	case res := <-ch1:
	// 		fmt.Println(res)
	// 	default:
	// 		fmt.Println("Timeout")
	// 	}
}
