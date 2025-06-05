package main

import "fmt"

type breed string

const (
	Shepherd breed = "овчарка"
	Corgi    breed = "corgi"
	Haska    breed = "haski"
)

type dog struct {
	Name   string
	breed  breed
	isTail bool
}

func NewDog(name string, breed breed, isTail bool) dog {
	return dog{
		Name:   name,
		breed:  breed,
		isTail: isTail,
	}
}

func (d dog) GetDogName() string {
	return d.Name
}

func setDogName(d *dog, newName string) {
	d.Name = newName
	fmt.Println("New name ", d.Name)
}

//Reciever - получатель
