package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AngkaMunculSekali(angka string) []int {
	pisah := strings.Split(angka, "")

	for x := 0; x < len(pisah); x++ {
		for y := x + 1; y < len(pisah); {
			if pisah[x] == pisah[y] {
				pisah = pisah[:y+copy(pisah[y:], pisah[y+1:])]
				pisah = pisah[:y+copy(pisah[x:], pisah[x+1:])]
				y = x + 1
			} else {
				y++
			}

		}
	}
	var hasil []int
	for z := 0; z < len(pisah); z++ {
		ambilBilangan := pisah[z]
		convert, err := strconv.Atoi(ambilBilangan)

		if err == nil {
			hasil = append(hasil, convert)
		}
	}
	return hasil

}

func main() {
	fmt.Println(AngkaMunculSekali("567895678"))
}
