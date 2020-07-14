package main

import "fmt"

func main() {
	perfect(5)
}

func perfect(iter int) {

	i := 0
	j := 0
	sum := 0

	for i = 1; i <= iter; i++ {
		sum = 0
		for j = 1; j < i; j++ {
			if i%j == 0 {
				sum += j
			}
		}
		if sum == i {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")
}
