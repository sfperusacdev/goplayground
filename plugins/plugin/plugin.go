package main

import "fmt"

var numcalls = 0

func EstaEsUnaFuncion(s string) {
	fmt.Println(numcalls, ":Hola como", s)
	numcalls++
}
