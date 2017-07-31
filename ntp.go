package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	now := time.Now()
	fmt.Printf("System time is %v\n", now)

	ntpTime, err := ntp.Time("ntp.local")
	if err != nil {
		fmt.Printf("Error\n")
	}
	fmt.Printf("NTP time is %v\n", ntpTime)

	difference := now.Unix() - ntpTime.Unix()
	fmt.Printf("Difference is %d seconds\n", difference)
}
