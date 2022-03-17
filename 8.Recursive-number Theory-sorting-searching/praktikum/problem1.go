package main

import (
	"fmt"
)

func primeX(num1, num2 int) {
	if num1 < 2 || num2 < 2 {
		fmt.Println("Numbers must be greater than 2.")
		return
	}
	for num1 <= num2 {
		isPrime := true
		for i := 2; i <= int(math.qrt(float64(num1))); i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d", num1)
		}
		num1++
	}
}
func main() {
	fmt.Println(primex(1))
	fmt.Println(primex(5))
	fmt.Println(primex(8))
	fmt.Println(primex(9))
	fmt.Println(primex(10))
}
