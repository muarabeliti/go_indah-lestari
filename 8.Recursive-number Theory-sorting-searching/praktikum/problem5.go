package main

import "fmt"

func FindMinAndMax(arr []int) string){
   fmt.Printf("Maximum element in %d, and %d is: %d\n", x, y, x - ((x - y) &
  	 ((x - y) >> 31)))
   fmt.Printf("Minimum element in %d, and %d is: %d\n", x, y, y + ((x - y) &
   	((x - y) >> 31)))
}
func main(){
   fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
   fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
   fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
   fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
   fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -28}))
}
