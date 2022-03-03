package main

import "fmt"

func playWithAsterik(n int) {
	//variabel
	var asterik string
	var star string = " * "

	//code
	for i := 1; 1 <= n; i++ {

		for j := n - 1; j >= 1; j-- {
		}
		asterik += star
		fmt.Println(asterik)
	}
}

func main() {
	playWithAsterik(5)
	/*

			   *
			 *  *
		    *  *  *
		   *  *  * *
		  * *  * *  *

	*/
}
