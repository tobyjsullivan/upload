package main

import (
	"net"
	"os"
	"fmt"
	"time"
	"log"
	"math/rand"
)

const (
	delay     = 1 * time.Millisecond
	maxMsgLen = 9216
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
	i := 0
	for {
		l := rand.Intn(maxMsgLen)
		message := make([]byte, l)
		_, err := rand.Read(message[:])
		if err != nil {
			panic(err.Error())
		}
		n, err := conn.Write(message[:])
		if err != nil {
			panic(err)
		}
		count += n
		i++

		if i % 100 == 0 {
			println("Sent:", count, "bytes")
		}

		time.Sleep(delay)
	}
}
