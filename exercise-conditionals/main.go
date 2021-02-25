package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
	year := 1993
	for year <= 2020 {
		fmt.Println(year)
		year++
	}

	year = 1993
	for {
		if year <= 2020 {
			break
		}
		fmt.Println(year)
		year++

	}

	for i := 10; i <= 100; i++ {
		fmt.Printf("When I divide %v between 4 I got the module %v\n", i/4, i%4)
	}

	if x := 2; x == 2 {
		fmt.Println("x is 2")
	} else if x == 5 {
		fmt.Println("x is 5")
	} else {
		fmt.Println("Anything")
	}

	switch {
	case 2 == 2:
		fmt.Println("2 is true")
		fallthrough
	case 2 == 3:
		fmt.Println("3 is true")
	}

	switch favSport := "football"; favSport {
	case "football":
		fmt.Println("It's football")
	case "basketball":
		fmt.Println("It's basketball")

	}
}
