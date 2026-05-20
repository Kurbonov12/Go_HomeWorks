package main

import "fmt"

func main() {
	a := 12
	b := 15
	c := 18
	V := a * b * c
	S := 2 * (a * b + b * c + a * c)
	fmt.Println(V)
	fmt.Println(S)
}