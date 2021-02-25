package main

import "fmt"

func main() {
	s1 := "Hello world"
	s2 := `This is a break 
	line`

	fmt.Println(s1)
	fmt.Printf("The type of s1 is: %T\n", s1)

	fmt.Println(s2)
	fmt.Printf("The type of s1 is: %T\n", s2)

	bs := []byte(s1)
	fmt.Println(bs)
	fmt.Printf("The type of s1 is: %T\n", bs)

	fmt.Println("")
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U ", s1[i])
	}
	fmt.Println("")

	for i, v := range s1 {
		fmt.Printf("At the index %d the value es %#X\n", i, v)
	}
	fmt.Printf("With the verb %q I tell the compiler to print the string %s", "%s", s1)

}
