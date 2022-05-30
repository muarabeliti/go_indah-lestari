package main

import (
	"fmt"
	"os"
	"strconv"
)

func getRootVariant(a, b, c int) (x, y, z int) {
	for x <= a || x <= b || x <= c {
		y = 0
		for y <= a || y <=b || y <= c {
			z = 0
			for z <= a || z <= b || z <= c {
				isMatchOnEquationA := x+y+z == a
				isMatchOnEquationB := x+y+z == b
				isMatchOnEquationC := x+x+y+y+z+z == c
				if isMatchOnEquationA && isMatchOnEquationB && isMatchOnEquationC {
					return
				}
				z++
			}
			y++
		}
		x++
	}
	panic("not solved")

	// test 1: 6 6 14 -> 1 2 3
	// test 2 : 1 2 3 -> panic
	func main () {
		var err error 
		var arguments = os.Args[1:]
		var numbers = make([]int, len(arguments))

		for index, number := range numbers {
			if number, err = strconv.Atoi(arguments[index]); err != nil {
				panic(err)
			}
			numbers[index]= number
		}
	}
}