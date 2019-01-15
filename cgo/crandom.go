package main

import "fmt"

/*
#include <stdlib.h>
*/
import "C"

func random() int {
	return int(C.random())
}

func main() {
	fmt.Println(random())
}
