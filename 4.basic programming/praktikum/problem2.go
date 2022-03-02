package main

import "fmt"

func main() {
	studentScore := 90

	if studentScore > 90 {
		fmt.Println("A")
	} else if studentScore < 85 {
		fmt.Println("B")
	} else if studentScore < 70 {
		fmt.Println("C")
	} else if studentScore < 60 {
		fmt.Println("D")
	} else if studentScore < 40 {
		fmt.Println("E")
	} else {
		fmt.Println("nilai invalid")
	}
}
