package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Println("error discovering port:", err)
		return
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("error accepting connection:", err)
			return
		} else {
			fmt.Println("connection has been accepted", c)
		}

		for i := range getLinesChannel(c) {
			fmt.Printf("read: %s\n", i)
		}

		fmt.Println("connection has been closed")
	}
}

func getLinesChannel(c net.Conn) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		sliceOfMessages := make([]byte, 2048)

		currentLine := ""

		sliceOfString := make([]string, 2)

		defer c.Close()

		for {
			count, err := c.Read(sliceOfMessages)
			if err != nil {
				if errors.Is(err, io.EOF) {
					if len(currentLine) > 0 {
						ch <- currentLine
					}

					break
				}
				fmt.Println("error reading connection:", err)
				return
			}

			sliceOfString = strings.Split(string(sliceOfMessages[:count]), "\n")

			for i := range len(sliceOfString) - 1 {
				currentLine += sliceOfString[i]

				ch <- currentLine

				currentLine = ""
			}
			currentLine += sliceOfString[len(sliceOfString)-1]
		}

	}()

	return ch
}
