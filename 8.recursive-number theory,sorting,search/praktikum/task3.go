package main

import (
  "fmt"
  "sort"
)

type stones []int

func (stone *stones) pop() int {
  popInt := (*stone)[len(*stone)-1]
  newStone := (*stone)[:len(*stone)-1]
  *stone = newStone
  return popInt
}

func (stone *stones) push(pushInt int) {
  var newStone stones = append(*stone, pushInt)
  *stone = newStone
}

func main() {
  var weights stones = []int{2, 4, 5, 9, 4, 10}

  for len(weights) > 1 {
    sort.Slice(weights, func(i, j int) bool {
      return weights[i] < weights[j]
    })
    a := weights.pop()
    b := weights.pop()
    if a != b {
      weights.push(a - b)
    }
    fmt.Println(weights)
  }

  if len(weights) > 0 {
    fmt.Print(weights[0])
  } else {
    fmt.Print("0")
  }

}

func main() {
	PrimaSegiEmpat(2, 3, 13)
	/*
	 17 19
	 23 29
	 31 37
	 156
	 */
	 primaSegiEmpat(5, 2, 1)
	 /*
	 2 3 5 7 11
	 13 17 19 23 29
	 129
	 */
}