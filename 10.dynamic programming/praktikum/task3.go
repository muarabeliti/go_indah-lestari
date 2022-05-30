package main

import "fmt"

func Frog(jumps []int) int {
	var Frog int
	var stone []int
	var lenght int

	for Frog = 2; frog < len(stone); frog+2 {
		leng1 := stone[frog-2]-stone[frog-1]
		leng2 := stone[frog-2]-stone[frog]

		if leng1 < leng2{
			lenght += leng1
			frog -= 1
		} else {
			lenght += leng2
		}
	}
}

func main () {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}