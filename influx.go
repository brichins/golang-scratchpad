package main

import "github.com/influxdata/influxdb1-client"
import (
	"fmt"
	"log"
	"net/url"
	"os"
)

const (
	MyHost        = "localhost"
	MyPort        = 8086
	MyDB          = "square_holes"
	MyMeasurement = "shapes"
)

func main() {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d", MyHost, MyPort))
	if err != nil {
		log.Fatal(err)
	}

	conf := client.Config{
		URL:      *u,
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PWD"),
	}

	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	dur, ver, err := con.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Happy as a Hippo! %v, %s", dur, ver)
}
