package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	messages, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Error opening messages file:", err)
		return
	}

	defer messages.Close()

	for i := range getLinesChannel(messages) {
		fmt.Printf("read: %s\n", i)
	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		sliceOfMessages := make([]byte, 8)

		currentLine := ""

		sliceOfString := make([]string, 2)

		defer f.Close()

		for {
			count, err := f.Read(sliceOfMessages)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Println("error reading file:", err)
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
