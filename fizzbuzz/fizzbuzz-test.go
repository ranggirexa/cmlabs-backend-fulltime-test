package main

import (
	"fmt"
)

func main() {
	// flag.Parse()
	// s := flag.Arg(0)
	for i := 0; i < 100; i++ {

		if (i+1)%3 == 0 && (i+1)%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if (i+1)%3 == 0 {
			fmt.Println("fizz")
		} else if (i+1)%5 == 0 {
			fmt.Println("buzz")
		}
	}
}
