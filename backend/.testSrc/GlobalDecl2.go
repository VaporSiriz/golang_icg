package main

import (
	"fmt"
)

type GLobal struct {
	global string
}

var Global GLobal
var global string

func main() {
	var q GLobal
	global = "3"
	q.global = "1"
	global2 = "5"
	var p GLobal
	p.global = "4"
	fmt.Printf("global : %s %s %v %v\n", global, global2, q, p)
	main2()
}

func main2() {
	
}

var global2 string