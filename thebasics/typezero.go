package main

import "fmt"

var y int
var x string
var z bool

func main() {
			// DECLARE a variable to be of a certain TYPE
			// then ASSIGN a VALUE of that type to the variable

			// Use VAR for zero value or package level scope

			fmt.Printf("%T\n", y)
			fmt.Println(y)
			fmt.Printf("%T\n", x)
			fmt.Println(`"`,x,`"`)
			fmt.Printf("%T\n", z)
			fmt.Println(z)

}