package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	a, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		fmt.Println("Error resolving address", err)
	}

	c, err := net.DialUDP("udp", nil, a)
	if err != nil {
		fmt.Println("Error connecting to address", err)
	}

	defer c.Close()

	b := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		messages, err := b.ReadString('\n')
		if err != nil {
			log.Println("Error reading data", err)
		}

		count, err := c.Write([]byte(messages))
		if err != nil {
			log.Println("Error writing data to UDP", err, count)
		}
	}
}
