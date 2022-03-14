package main

import (
	"fmt"
)

func Caesar(offset int, input string) string {
	var output string
	var temp rune

	for _, value := range input {
		//bila lowercase
		temp = ((value + rune(offset) - 97) % 26) + 97
		//output
		output += string(temp)
	}

	return output
}

func main() {
	fmt.Println(Caesar(3, "abc"))
	fmt.Println(Caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(Caesar(10, "alteraacademy"))
	fmt.Println(Caesar(1, "abcdefghijklmnopqrstuvwxyz"))

}
