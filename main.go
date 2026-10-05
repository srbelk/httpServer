package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	messages, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Error opening messages file:", err)
		return
	}

	defer messages.Close()

	sliceOfMessages := make([]byte, 8)

	for {
		count, err := messages.Read(sliceOfMessages)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Println("error reading file:", err)
			return
		}

		fmt.Printf("read: %s\n", sliceOfMessages[:count])
	}
}
