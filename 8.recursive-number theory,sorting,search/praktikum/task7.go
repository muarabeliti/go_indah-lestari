package main

import "fmt"

func playingDomino(card[][]int, deck []int) interface {} {
	var cardsame [][]int
	for cardNumber := 0; cardNumber < len(card); cardNumber++ {
		if card[cardNumber][0] == deck[0] || card[cardNumber] [1] == deck[1] ||
			card[cardNumber][0] == deck[1] || card[cardNumber] [1] == deck[0] {
				cardsame = append(cardsame, card[cardNumber])
			}
	}

	if len(cardsame)>= 2 {
		var max, index int = 0, 0
		for cardSameNuber := 0; cardSameNuber < len(cardsame); cardSameNuber++ {
			total := cardsame[cardSameNuber][0]+ cardsame[cardSameNuber][1]
			if max <= total {
				max = total
				index = cardSameNuber
			}
		}

		return  cardsame[index]
	} else if len(cardsame) >= 1 {
		return cardsame[0]
	} else {
		return "tutup kartu"
		}

	}

	func main() {
		fmt.Println(playingDomino([][]int {[]int {6, 5}, []int {3, 4}, []int {2, 1}, []int{3, 3}, []int {4,3}))
		fmt.Println(playingDomino([][]int {[]int {6, 5}, []int {3, 3}, []int {3, 4}, []int{2, 1}, []int {3,6}))
		fmt.Println(playingDomino([][]int {[]int {6, 6}, []int {2, 4}, []int {3, 6}, []int{5, 1}))
	}

}