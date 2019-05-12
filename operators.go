package main
import "fmt"

// v := 'hi'



func main() {
	// := declares and assigns a value to variable
	// you can reassign value with just an = sign later


	y := []string{"1", "2", "3"}
	fmt.Println(y)
	// fmt.Println(v)
	a, b, c, d := f2()
	fmt.Println(d)
	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(a)
}

func f2() (int, int, int, int) {
	return 3, 4, 5, 6 
}
// func foo(x int) {
// 	fmt.Println(x)
// }