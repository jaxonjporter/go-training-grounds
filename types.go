package main

import (
	"fmt"
)
type Dog struct {
	Animal
}
type Animal struct {
	Age int
}

func (a *Animal) Move() {
	fmt.Println("Animal moved")
}
func (a *Animal) SayAge() {
	fmt.Printf("Animal age: %d\n", a.Age)
}
func (a *Animal) Eat() {
	fmt.Printf("You fed the dog.\n")
}
func main() {
	d := Dog{}
	d.Age = 3
	d.Move()
	d.SayAge()
	d.Eat()
}