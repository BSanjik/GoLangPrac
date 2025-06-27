package main

import "fmt"

func goroutineOne(ch chan int) {
	result := <-ch
	fmt.Println(result)
}

func main() {

	// myChan := make(chan int)

	// go goroutineOne(myChan)

	// myChan <- 42

	// fmt.Println("hello!")

	// myChan := make(chan string, 1)

	// myChan <- "Almat champ of Hackathon"

	// fmt.Println(<-myChan)

	//task 1
	// chan1 := make(chan string, 1)
	// chan1 <- "Hello world"
	// fmt.Println(<-chan1)

	//task 2
	// chan2 := make(chan []int, 1)
	// chan2 <- []int{1, 2, 3, 4}
	// slice := <-chan2
	// for i := range slice {
	// 	fmt.Println(slice[i])
	// }

	//task 3
	// chan3 := make(chan []int, 1)
	// chan3 <- []int{0, 1, 2, 3, 4, 5}
	// close(chan3)
	// fmt.Println(<-chan3)
	// _, ok := <-chan3
	// fmt.Println(ok)
	// if ok {
	// 	fmt.Println(<-chan3)
	// 	close(chan3)
	// }
	// fmt.Println(<-chan3)
	// fmt.Println(<-chan3)
	// fmt.Println(<-chan3)
	// //result, ok := <-chan3
	// // fmt.Println(result, ok)

	//task 4

}
