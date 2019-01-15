package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, _ := plugin.Open("hexify.so")
	f, _ := p.Lookup("Hexify")
	fmt.Println(f.(func(string) string)("gopher"))
	// 676f70686572
}
