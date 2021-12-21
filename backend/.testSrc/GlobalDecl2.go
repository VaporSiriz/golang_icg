package main

import (
	"fmt"
)

type Global struct {
	global string
}

var global string

func main() {
	var q Global
	global = "3"
	q.global = "1"
	fmt.Printf("global : %s %v\n", global, q)
}