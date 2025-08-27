package main

import "fmt"

func main() {

	// In Go functions are treated as first-class citizens , means 
	// in a given programming language design, a first-class citizen[a] is an entity which supports all the operations generally available to other entities. These operations typically include being passed as an argument, returned from a function, and assigned to a variable.

	// Expressions -> It is a combination of variables, operators and function calls that evaluates to a single value.



	x := foo()

	x1 := foo

	x1()

	fmt.Println(x)

	y := bar()

	y1 := bar

	y1()

	fmt.Println(y())

	z := cat()

	z()

	z1 := cat

	z1()

	fmt.Printf("%T\t%T\t%T\n",x,y,z)

	fmt.Printf("%T\t%T\t%T\n", foo, bar, y)

	fmt.Printf("%T\t%T\t%T\n", x1,y1,z1)
}

func foo() int{
	return 42
}

func bar() func() int {

	return func() int {
		return 32
	}
}

func cat() func() {
	return func() {
		fmt.Println("I cat!")
	}
}
