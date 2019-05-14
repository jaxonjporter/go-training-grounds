package main

import (
	"fmt"
)
type Dog struct {
	Animal
}
type Cat struct {
	Animal
}
type Lizard struct {
	Animal
}
type Snake struct {
	Animal
}
type Animal struct {
	Age int
	Color string

}

func (a *Animal) Move() {
	fmt.Println("Animal moved")
}
func (a *Animal) SayFur() {
	fmt.Printf("Animal Fur Color: %s\n", a.Color)
}
func (a *Animal) SayAge() {
	fmt.Printf("Animal age: %d\n", a.Age)
}
func (a *Animal) Eat() {
	fmt.Printf("You fed the animal.\n")
}
func main() {
	
	d := Dog{}

	d.Age = 3
	d.Color = "White"
	d.SayFur()
	d.Move()
	d.SayAge()
	d.Eat()
	
	c := Cat{}
	
	c.Age = 3
	c.Color = "Grey"
	c.SayFur()
	c.Move()
	c.SayAge()
	c.Eat()

	l := Lizard{}

	l.Age = 6
	l.Color = "None"
	l.SayFur()
	l.Move()
	l.SayAge()
	l.Eat()

	s := Snake{}

	s.Age = 20
	s.Color = "None"
	s.SayFur()
	s.Move()
	s.SayAge()
	s.Eat()
	
}