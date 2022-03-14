package main

import (
	"fmt"
	"strings"
)

func compare(a, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}
	cekA := strings.Contains(a, b)

	if cekA {
		return b
	} else {
		return a

	}

}

func main() {
	fmt.Println(compare("AKA", "AKASHI"))
	fmt.Println(compare("KANGOORO", "KANG"))
	fmt.Println(compare("KI", "KIJANG"))
	fmt.Println(compare("KUPU-KUPU", "KUPU"))
	fmt.Println(compare("ILALANG", "ILA"))
}
