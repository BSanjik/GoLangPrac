package main

import (
	"fmt"
	"time"
)

type Number int

type Event struct {
	name string
	date time.Time
}

func main() {
	// // Replace the following line with a correct call to NewDog if it's in the same package
	// myDog := NewDog("Hatiko", Corgi, true)

	// fmt.Println(myDog.Name)
	// setDogName(&myDog, "Mops")
	// fmt.Println(myDog.GetDogName())
	// // anotherBreed := "Mops-Alabai"
	// // fmt.Println(NewDog("Hatiko", "Akito", true))
	// // fmt.Println(NewDog("Hatiko", breed(anotherBreed), true))
	// // //fmt.Println(NewDog("Hatiko", corgi, true))
	fmt.Println("Write task number")
	var task int
	fmt.Scan(&task)
	switch task {
	case -1:
		lessontask1()
	case 1:
		task1()
	case 2:
		task2()
	case 3:
		task3()
	}
}

func lessontask1() {
	var eventName string
	var eventDateStr string
	fmt.Print("Enter event name: ")
	fmt.Scan(&eventName)
	fmt.Print("Enter event date (YYYY-MM-DD): ")
	fmt.Scan(&eventDateStr)
	eventDate, err := time.Parse("2006-01-02", eventDateStr)
	if err != nil {
		fmt.Println("Invalid date format")
		return
	}
	event := newEvent(eventName, eventDate)
	fmt.Printf("Дней до события: %d\n", event.DaysUntil())
}

func newEvent(eventName string, date time.Time) Event {
	return Event{
		name: eventName,
		date: date,
	}
}

func (event Event) DaysUntil() int {
	today := time.Now().Truncate(24 * time.Hour)
	eventDay := event.date.Truncate(24 * time.Hour)
	diff := eventDay.Sub(today).Hours() / 24
	if diff < 0 {
		return 0
	}
	return int(diff)
}

func task1() {
	var num Number
	fmt.Scan(&num)
	if num.task1Even() {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is not even")
	}
}

func (num Number) task1Even() bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func task2() {
	var num Number
	fmt.Scan(&num)
	if num.task2Prime() {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is not even")
	}
}

func (num Number) task2Prime() bool {
	if num < 2 {
		return false
	}
	for i := 2; Number(i*i) <= num; i++ {
		if num%Number(i) == 0 {
			return false
		}
	}
	return true
}

func task3() {
	var n int
	var number Number
	fmt.Scan(&n)
	slice := make([]Number, n)
	for i := range n {
		fmt.Scan(&slice[i])
	}
	fmt.Println(number.task3Slice(slice))
}

func (num Number) task3Slice(slice []Number) int {
	var result int
	for i := range slice {
		result += int(slice[i])
	}
	return result
}
