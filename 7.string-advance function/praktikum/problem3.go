package main

import (
	"fmt"
)

func swap(a, b *int) {
	var temp = *a
	*a = *b
	*b = temp
}
func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Printf("a : %d\t b : %d", a, b)
}
