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

	sliceOfMessages := make([]byte, 8)

	currentLine := ""

	sliceOfString := make([]string, 2)

	for {
		count, err := messages.Read(sliceOfMessages)
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

			fmt.Printf("read: %s\n", currentLine)

			currentLine = ""
		}

		currentLine += sliceOfString[len(sliceOfString)-1]
	}
}
