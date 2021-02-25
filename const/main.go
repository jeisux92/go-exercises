package main

import "fmt"

const a = 42
const b = 42.32
const c = "Edwar Tua"
const f, e = 32, 12
const (
	t  = 32
	td = 12
)

type name string

var nn name

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	nn = c

}
