package main

import (
	"fmt"
)

var a int
var b string = "program"
var c bool
var d = true

var x, f = 42, true

type money int

var m money

func main() {
	e := 42
	f := "says hello world"
	g := `The doctor says that eating vegetables is 
healthy`
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println("My", b, f)
	fmt.Println(g)
	fmt.Printf("the value of the variable a is: %#v\n", a)
	fmt.Printf("the value of the variable b is: %v\n", b)
	fmt.Printf("the type of a is: %T\n", a)
	fmt.Printf("the type of b is: %T\n", b)

	s1 := fmt.Sprint("The ", b, " says hello world")
	fmt.Println(s1)

	fmt.Println(m)
	fmt.Printf("the type of m is: %T\n", m)
	m = 100
	converted := int(m)
	fmt.Printf("the value of the variable converted is: %v\n", converted)
}
