package main
import "fmt"

// v := 'hi'


// Go is a static programming language. When assigned a value, that 
// variable is designed to hold a value of a specific type.

var x = 3

// var makes it globally scoped

func main() {
	// := declares and assigns a value to variable
	// you can reassign value with just an = sign later


	// backticks are raw string litterals
	fmt.Println(`again`, x)
	y := []string{"1", "2", "3"}
	fmt.Println(y)
	// fmt.Println(v)
	a, b, c, d := f2()
	fmt.Println(d)
	fmt.Println(c)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(a)
}

func f2() (int, int, int, int) {
	return 3, 4, 5, 6 
}
// func foo(x int) {
// 	fmt.Println(x)
// }


// use short declaration as much as possible