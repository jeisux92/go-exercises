package main

import "fmt"

func main() {
	switch "B" {
	case "A", "B":
		fmt.Print("It's A")
		fallthrough
	default:
		fmt.Print("It's anything")

	}
}
