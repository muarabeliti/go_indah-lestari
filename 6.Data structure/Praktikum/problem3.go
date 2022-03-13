package main

import (
	"fmt"
	"sort"
)

func pairSum(arr []int, target int) []int {
	data := []int{}
	sort.Ints(arr)
	awal, akhir := 0, len(arr)-1
	for awal < akhir {
		//awal + akhir
		total := arr[awal] + arr[akhir]
		//jika true masukan ke slice data
		if total == target {
			data = append(data, awal, akhir)
			return data
			//jika tidak pair di awal maka awal ++
		} else if total < target {
			awal++
			//jika awal tidak pair pindah ke akhir --
		} else {
			akhir--
		}
	}
	return data
}

func main() {
	fmt.Println(pairSum([]int{1, 2, 3, 4, 5, 6}, 6))
	fmt.Println(pairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(pairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(pairSum([]int{1, 4, 6, 18}, 10))
	fmt.Println(pairSum([]int{1, 5, 6, 7}, 6))
}
