package main

import "fmt"

func ArrayMarge(ArrayA, ArrayB []string) []string {
	var hasil []string
	hasil = append(hasil, ArrayA...)
	hasil = append(hasil, ArrayB...)

	for x := 0; x < len(hasil); x++ {
		for y := x + 1; y < len(hasil); {
			if hasil[x] == hasil[y] {
				hasil = hasil[:y+copy(hasil[y:], hasil[y+1:])]
				y += 0
			} else {
				y++
			}
		}
	}

	return hasil

}

func main() {
	fmt.Println(ArrayMarge([]string{"king", "devil jin", "eddie"}, []string{"eddie", "steve", "gees"}))
}
