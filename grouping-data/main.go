package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", a)

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", x)
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	// state
	states := make([]string, 50, 50)
	states = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(len(states))
	fmt.Println(cap(states))
	for i := 0; i < len(states); i++ {

		fmt.Println(i, states[i])
	}

	// nested array
	batman := []string{"Batman", "jefe", "Vestido Negro"}
	robin := []string{"Robin", "ayudante", "Vestido de Colores"}

	superHeroes := [][]string{batman, robin}

	fmt.Println(superHeroes)

	for i, reg := range superHeroes {
		fmt.Println("Register: ", i)
		for j, val := range reg {
			fmt.Printf("\tIndex of position: %v\tvalue: %v\n", j, val)
		}
	}

	// maps
	people := map[string][]string{
		`eduar_tua`:   {"computers", "montains", "beach"},
		`gabriel_mar`: {"reading", "buying", "music"},
		`ana_fan`:     {"icecream", "painting", "dancing"},
	}

	people["gloria_bust"] = []string{"icecream", "painting", "dancing"}

	if _, ok := people["ana_fan"]; ok {
		delete(people, `ana_fan`)
	}

	for k, v := range people {
		fmt.Println("Register: ", k)
		for j, val := range v {
			fmt.Printf("\tIndex of position: %v\tvalue: %v\n", j, val)
		}
	}
}
