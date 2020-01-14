package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Input: ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			break
		}
	}

	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "%s", scanner.Err().Error())
	}
}
