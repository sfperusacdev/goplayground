package main

import "fmt"

func main() {
	var bite byte = 61
	fmt.Printf("%b\n", bite>>2&1)
}
