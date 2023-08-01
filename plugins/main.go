package main

import (
	"fmt"
	"log"
	"plugin"
)

func main() {
	mod, err := plugin.Open("miplugin.so")
	if err != nil {
		log.Fatalln(err)
	}
	symbol, err := mod.Lookup("EstaEsUnaFuncion")
	if err != nil {
		log.Fatalln(err)
	}

	fn, ok := symbol.(func(s string))
	if !ok {
		fmt.Println("unexpected type from module")
		return
	}
	fn("Kevin")
	fn("Kevin")
	fn("Kevin")
	fn("Kevin")
}
