package main

import "fmt"

var a int
type hotdog int
var b hotdog

func main() {
			a = 42
			b = 45
			fmt.Println(a)
			fmt.Printf("%T\n", a)

			fmt.Println(b)
			fmt.Printf("%T\n", b)

			// you can convert (not called casting in go)

			a = int(b)
			// a is now the value of b, as type int
			fmt.Println(a)
			fmt.Printf("%T\n", a)

			b = hotdog(a)
			//  b is now the value of a, as type hotdog
			fmt.Println(b)
			fmt.Printf("%T\n", b)
}