package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var ran int

	rand.Seed(1)
	ran = rand.Intn(10)

	fmt.Println(ran)
}
