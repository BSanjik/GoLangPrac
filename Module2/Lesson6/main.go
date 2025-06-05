package main

import "fmt"

type Fish struct{}

func (f Fish) Swim() {
	fmt.Println("fish SwimS")
}

type Boat struct{}

func (b Boat) Swim() {
	fmt.Println("boat SwimS")
}

//Inteface
type Swimmer interface {
	//methods
	Swim()
}

func ObjectSwim(s Swimmer) {
	s.Swim()
}

type Cat struct{}

func (c Cat) Voice() {
	fmt.Println("Meow")
}

type Dog struct{}

func (d Dog) Voice() {
	fmt.Println("GAV")
}

type Speaker interface {
	Voice()
}

func GetVoice(s Speaker) {
	s.Voice()
}

//task1
type Message struct {
	name string
}

func (m Message) Print() {
	fmt.Println("Message is", m.name)
}

type Printer interface {
	Print()
}

func GetPrint(p Printer) {
	p.Print()
}

//task2
type MyString string

type Stringer interface {
	Length() int
}

func (m MyString) Length() int {
	return len(string(m))
}

func GetLength(s Stringer) {
	s.Length()
}

func main() {
	// fish := Fish{}
	// boat := Boat{}
	// cat := Cat{}
	// dog := Dog{}

	//task1
	// message := Message{}

	// fmt.Scan(&message.name)

	// GetPrint(message)

	//task2
	// var MyString MyString
	// fmt.Scan(&MyString)
	// fmt.Println(MyString.Length())

	// rectangle := Rectangle{10, 5}
	// circle := Circle{5}

	// fish.FishSwim()
	// boat.BoatSwim()
	// ObjectSwim(fish)
	// ObjectSwim(boat)

	// GetVoice(cat)
	// GetVoice(dog)

	// getArea(rectangle)
	// getArea(circle)

	warrior := Warrior{
		Character: Character{
			Name: "Vitya",
			HP:   100,
			Dps:  10,
		},
		WeaponDMG: 10,
	}
	dragon := Dragon{
		Character: Character{
			Name: "Smaug",
			HP:   1000,
			Dps:  10,
		},
	}

	getDamage(warrior, &dragon.Character)
	fmt.Println(dragon.HP)

	getDamage(dragon, &warrior.Character)
	fmt.Println(warrior.HP)
}

type Character struct {
	Name string
	HP   int
	Dps  int
}

type Warrior struct {
	Character
	WeaponDMG int
}

func (w Warrior) Attack(c *Character) {
	c.HP = c.HP - (w.Dps + w.WeaponDMG)
}

type Dragon struct {
	Character
}

func (d Dragon) Attack(c *Character) {
	c.HP = c.HP - d.Dps
}

type Attacker interface {
	Attack(*Character)
}

func getDamage(a Attacker, wOrd *Character) {
	a.Attack(wOrd)
}

type Rectangle struct {
	Width, Height int
}

func (r Rectangle) Area() float64 {
	return float64(r.Width * r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func getArea(s Shape) {
	fmt.Println(s.Area())
}

type Shape interface {
	Area() float64
}
