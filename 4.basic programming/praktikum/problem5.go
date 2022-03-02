package main

import "fmt"

func palindrome (kata string) {
	var hasilCek,hasilAkhir string
	var panjangHuruf int
	panjangHuruf = len(kata)

	for x :=0; x < int(panjang Huruf) ; x++ {
		y := (panjangHuruf -1)-x
		hasilCek = fmt.Sprintf("%c",kata[y])

	if hasilCek != fmt.Sprintf("%c", kata[x]){
		hasilAkhir ="bukan palindrome"
		break
	}else {
		hasilAkhir ="palindrome"
	}


	}
fmt.Println("kata",kata,"merupakan=",hasilAkhir)

}

func main (){
	palindrome ("civic")
	palindrome ("katak")
	palindrome ("kasur rusak")
	palindrome ("mistar")
	palindrome ("lion")
}