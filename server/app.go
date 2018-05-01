package main

import (
	"os"
	"log"
	"net"
	"fmt"
)

const (
	message = "Hello, world!"
)

func main()  {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatalln("Must define PORT")
	}

	conn, err := net.ListenPacket("udp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err.Error())
	}

	count := 0
	i := 0
	buffer := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err.Error())
		}

		//content := buffer[:n]
		count += n
		i++

		if i % 100 == 0 {
			println("Received:", count, "bytes")
		}
	}
}
