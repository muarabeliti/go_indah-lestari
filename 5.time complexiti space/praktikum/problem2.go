package main

import "fmt"

func powerify(base int, power int) int {
	if power == 0 {
		return 1
	}
	var result int

	result = powerify(base, power/2)
	result = result * result

	if power%2 == 1 {
		result = result * base
	}

	return result
}

func main() {
	fmt.Println(powerify(5, 3))
}
