package main

import "fmt"

var u int
var v string
var w bool

type myType int

var mt myType
var mySub int

func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
	fmt.Println(u)
	fmt.Println(v)
	fmt.Println(w)

	fmt.Printf("Hello\n")
	fmt.Printf("\tHello\n")
	fmt.Printf("\t\ttHello\n")

	fmt.Println("the value of my new type is:", mt)
	fmt.Printf("the type of nt is: %T", mt)
	mt = 42
	fmt.Println("the new value of my new type is:", mt)
	mySub = int(mt)
	fmt.Println("the value of mySub is:", mySub)
	fmt.Printf("the type of mySub is: %T", mySub)

}
