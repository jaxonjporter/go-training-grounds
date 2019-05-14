package main

import "fmt"

var y = 987698

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\t%T\t%b\t%v\t%#q\n", y, y, y, y, y)
	s := fmt.Sprintf("%#x\t%T\t%b\t%v\t%#q\n", y, y, y, y, y)
	fmt.Println(s)
}