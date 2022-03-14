package main

import (
	"fmt"
)

func minmax(num ...*int) (min int, max int) {
	var k int
	min = *(num[0])
	for k = 0; k < len(num); k++ {
		if *(num[k]) > max {
			max = *(num[k])
		} else if *(num[k]) < min {
			min = *(num[k])
		}
	}
	return
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Println("Masukkan a1 : ")
	fmt.Scan(&a1)
	fmt.Println("Masukkan a2 : ")
	fmt.Scan(&a2)
	fmt.Println("Masukkan a3 : ")
	fmt.Scan(&a3)
	fmt.Println("Masukkan a4 : ")
	fmt.Scan(&a4)
	fmt.Println("Masukkan a5 : ")
	fmt.Scan(&a5)
	fmt.Println("Masukkan a6 : ")
	fmt.Scan(&a6)

	min, max = minmax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("nilai min : ", min)
	fmt.Println("nilai min : ", max)

}
