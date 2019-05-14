// you will always start with package main
// and exit with func main

package main
import "fmt"

func main() {
	fmt.Println("Hello Everyone!\n")
	foo()
	fmt.Println("something more")

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			n, _ := fmt.Println(i)
			fmt.Println(n)
			foo()
		}
	}
	bar()
}

func foo() {
	fmt.Println("Im in foo\n")
}

func bar() {
	fmt.Println("Exiting...")
}
// control flow:
// (1) sequence
// (2) loop: iterative
// (3) conditional

// Runs Top to Bottom

// undeline means ignore to compiler
//  Never have a variable you dont use