package main

import "fmt"

func gcd(a int, b int) int {
	c := a
	d := b

	if c == 0 {
		return d
	}
	for d != 0 {
		if c > d {
			c = c - d
		} else {
			d = d - c
		}
	}
	return c
}

func main() {
	a := 60
	b := 48
	gcd := gcd(a,b)
	fmt.Println(gcd)
}
