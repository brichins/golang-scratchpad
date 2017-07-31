package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	resp, err := conn.Do("PING")
	fmt.Println(resp)
}
