package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func isPalindrome(input string) bool {
	// fmt.Println(len([]rune(input)))

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func removeSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}

func main() {
	for {
		fmt.Printf("Masukan Kata : ")

		input := bufio.NewReader(os.Stdin)
		inputr, _ := input.ReadString('\n')
		str := removeSpace(inputr)

		result := isPalindrome(str)
		if result == true {
			fmt.Println("palindrome")
		} else {
			fmt.Println("not palindrome")
		}
	}
}
