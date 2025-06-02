package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type mile int
type color string

type Cat struct {
	Color  color `json: "superColor"`
	Breed  string
	Gender int
}

type Book struct {
	Name   string
	Author string
}
type Library struct {
	Books []Book
}

type Creature struct {
	Name string
	Age  int
}

type fruits struct {
	Name   string
	Color  string
	Weight int
}

type Person struct {
	Name string
	Age  int
}

func AddBook(library *Library) {
	book := Book{
		Name:   "Clean Code",
		Author: "Robert Martin",
	}
	library.Books = append(library.Books, book)
}

func printLibrary(library Library) {
	fmt.Println(library)
}

func setColor(f fruits) fruits {
	f.Color = "green"
	return f
}

// task 1
func getPersonInfo(personInfo Person) {
	fmt.Println("Name ", personInfo.Name, "Age ", personInfo.Age)
}

// Construction function
func NewCat(color color, breed string, gender int) Cat {
	return Cat{
		Color:  color,
		Breed:  breed,
		Gender: gender,
	}
}

func main() {
	cat1 := NewCat("brown", "Meinkun", 1)
	fmt.Println(cat1)
	//Marshal
	cat1Bytes, err := json.Marshal(cat1)
	if err != nil {
		log.Println("Error")
		return
	}
	fmt.Println(string(cat1Bytes))
	//Unmarshal
	var acceptedCat Cat
	err = json.Unmarshal(cat1Bytes, &acceptedCat)
	if err != nil {
		log.Println("Error", err)
		return
	}
	fmt.Println(acceptedCat)

	// var lib Library
	// AddBook(&lib)
	// AddBook(&lib)
	// printLibrary(lib)

	// var name string
	// var age int
	// p := Person{
	// 	Name: name,
	// 	Age:  age,
	// }
	// fmt.Scan(&p.Name, &p.Age)
	// getPersonInfo(p)
	// c := Creature{
	// 	Name: "Alisher and Shark",
	// 	Age:  25,
	// }
	// fmt.Println(c.Name)
	// fmt.Println(c.Age)
	// var one mile
	// one = 10
	// var two int = 5
	// fmt.Println(one + mile(two))
	// fmt.Println(one)
}

func lessonTask(distacne mile) string {
	return "Dis to the obj " + string(distacne)

}
