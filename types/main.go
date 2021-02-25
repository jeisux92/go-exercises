package main

import (
	"fmt"
	"runtime"
)

var s int8

func main() {
	x := 42.32
	y := 42
	s = -128

	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", s)
	fmt.Printf("type: %T value: %v\n", runtime.GOOS, runtime.GOOS)
	fmt.Printf("type: %T value: %v\n", runtime.GOARCH, runtime.GOARCH)
}
