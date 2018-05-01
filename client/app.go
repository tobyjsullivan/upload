package main

import (
	"net"
	"os"
	"fmt"
	"time"
	"log"
)

const (
	delay = 1 * time.Millisecond
	message = "Hello, world!"
)

func main() {
	ipAddress := os.Getenv("DEST_IP_ADDR")
	port := os.Getenv("DEST_PORT")

	if ipAddress == "" {
		log.Fatalln("Must define DEST_IP_ADDR")
	}
	if port == "" {
		log.Fatalln("Must define DEST_PORT")
	}

	conn, err := net.Dial("udp", fmt.Sprintf("%s:%s", ipAddress, port))
	if err != nil {
		panic(err.Error())
	}

	count := 0
	for {
		_, err = conn.Write([]byte(message))
		if err != nil {
			panic(err)
		}
		count++

		if count % 100 == 0 {
			println("Count:", count)
		}

		time.Sleep(delay)
	}
}
