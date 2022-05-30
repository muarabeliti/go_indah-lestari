package main

import (
	"fmt"
)

func fibonacci(number int) int {
	var terms, next int
	var previous, current int = 0, 1
	fmt.Printf("\nMasukkan jumlah terms: ")
	fmt.Scanf("%d", &terms)
	if terms < 2 {
		fmt.Printf("\nJumlah minimal term pada Fibonacci series adalah dua")
	} else {
		fmt.Printf("\n\nFibonacci series:\n\n")
		for i := 0; i < (terms); i++ {
			fmt.Printf("%d ", previous)
			next = previous + current
			previous = current
			current = next
		}
	}
	
	fmt.Printf("\n")
}

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(9))
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci(12))
	

}