package main

import (
	"net"
	"os"
	"fmt"
	"time"
	"log"
	"math/rand"
	"context"
)

const (
	delay     = 1 * time.Millisecond
	maxMsgLen = 50000
	nThreads = 5
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

	ctx := context.Background()

	countPipe := make(chan int, 20)
	for i := 0; i < nThreads; i++ {
		go runThread(ipAddress, port, countPipe)
	}

	go runReports(countPipe)

	<-ctx.Done()
}

func runThread(ipAddress, port string, countPipe chan int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ipAddress, port))
	if err != nil {
		panic(err.Error())
	}

	for {
		message := make([]byte, maxMsgLen)
		_, err := rand.Read(message[:])
		if err != nil {
			panic(err.Error())
		}
		n, err := conn.Write(message[:])
		if err != nil {
			panic(err)
		}

		countPipe<- n

		//time.Sleep(delay)
	}
}

func runReports(countPipe chan int) {
	start := time.Now()

	i := 0
	totalBytes := 0
	for n := range countPipe {
		i++
		totalBytes += n

		if i%100 == 0 {
			elapsed := time.Now().Sub(start)
			secondsElapsed := elapsed.Seconds()
			uploadRateKbps := float64(float64(totalBytes) / secondsElapsed) / 1024
			println(fmt.Sprintf("Rate: %.1f kbps\tSent: %d bytes in %.2f seconds", uploadRateKbps, totalBytes, secondsElapsed))
		}
	}
}
