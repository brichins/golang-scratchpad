package main

import "fmt"
import "code.google.com/p/leveldb-go/leveldb"

func main() {
	fmt.Println("Hello, world")
	db, _ := leveldb.Open("/tmp/leveldb", nil)
	defer db.Close()
}
