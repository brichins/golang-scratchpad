package main

import (
	"fmt"
	"log"
	"net"
)

// Pings a daytime TCP server
func main() {
	conn, err := net.Dial("tcp", "daytime.local:13")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil || n == 0 {
		conn.Close()
	}
	fmt.Printf("%s", buf)
}
