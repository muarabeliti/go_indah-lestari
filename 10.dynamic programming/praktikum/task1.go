package main

import "fmt"

func fibo(n int) int {
    a := 0
    b := 1
    // Iterate until desired position in sequence.
    for i := 0; i < n; i++ {
        // Use temporary variable to swap values.
        temp := a
        a = b
        b = temp + a
    }
    return a
}

func main() {
    // Display first 11 Fibo numbers.
    for n := 0; n < 11; n++ {
        // Compute number.
        result := fibo(n)
        fmt.Printf("Fibo %v = %v\n", n, result)
    }
}
	