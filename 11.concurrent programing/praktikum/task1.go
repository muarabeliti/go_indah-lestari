// Golang program to count the number of
// repeating words in given Golang String
package main

import (
	"fmt"
	"strings"
)

func repetition(st string) map[string]int {

	// using strings.Field Function
	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}

// Main function
func main() {
	input := "lorem ipsum dolor sit amet , consectetur adipiscing elit, " +
		"sed do eiusmod tempor incididunt ut labore et dolore magna aliqua "
	for index, element := range repetition(input) {
		fmt.Println(index, "=", element)
	}
}
