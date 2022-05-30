package main

import (
    "fmt"
    "sort"
)

func BinarySearch(array[]int, x int) {
    a := []int
    x := 

    i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
    if i < len(a) && a[i] == x {
        fmt.Printf("found %d at index %d in %v\n", x, i, a)
    } else {
        fmt.Printf("%d not found in %v\n", x, a)
    }
}

func main() {
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7,} 3)
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10,} 5)
	BinarySearch([]int{12, 15, 19, 24, 31, 53, 59, 60} 53)
	BinarySearch([]int{12, 15, 19, 24, 31, 53, 59, 60} 100)
}