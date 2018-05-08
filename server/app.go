package main

import (
	"os"
	"log"
	"net"
	"fmt"
	"time"
)

const (
	bufferSize = 32 * 1024
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatalln("Must define PORT")
	}

	conn, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err.Error())
	}

	countPipe := make(chan int, 5)
	go runReports(countPipe)
	for {
		conn, err := conn.Accept()
		if err != nil {
			panic(err.Error())
		}

		go handleConn(conn, countPipe)
	}

}

func handleConn(conn net.Conn, countPipe chan int) {
	buffer := make([]byte, bufferSize)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			println("Error:", err.Error())
			break
		}

		countPipe <- n
	}
}

func runReports(countPipe chan int) {
	start := time.Now()
	totalBytes := 0
	i := 0
	for n := range countPipe {
		i++
		totalBytes += n

		if i%10000 == 0 {
			elapsed := time.Now().Sub(start)
			secondsElapsed := elapsed.Seconds()
			uploadRateKbps := float64(float64(totalBytes)/secondsElapsed) / 1024
			println(fmt.Sprintf("Rate: %.1f kbps\tSent: %d bytes in %.2f seconds", uploadRateKbps, totalBytes, secondsElapsed))
			start = time.Now()
			totalBytes = 0
		}
	}
}
