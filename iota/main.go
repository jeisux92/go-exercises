package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	d = iota
)

// each time we type const the iota is reseted

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
}
