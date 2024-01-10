package main

import "fmt"

// learn  odd number dan even number
func main() {
	var (
		even []int
		odd  []int
	)
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}
	}
	fmt.Println("Even Number =", even)
	fmt.Println("odd Number =", odd)
}
