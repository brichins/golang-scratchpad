package main

import "fmt"
import "github.com/golang/leveldb"

func main() {
	fmt.Println("Hello, world")
	db, _ := leveldb.Open("/tmp/leveldb", nil)
	defer db.Close()
}
