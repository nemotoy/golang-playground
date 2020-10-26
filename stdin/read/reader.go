package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args
	flag := "line"
	if len(args) != 0 {
		flag = args[1]
	}

	const testdata = "title: description\naaaaa\nbbbbb\nccccc"
	fmt.Println("**************INPUT**************")
	fmt.Println(testdata)
	fmt.Println("*********************************")

	reader := bufio.NewReader(strings.NewReader(testdata))
	fmt.Println("**************START**************")

	switch flag {
	case "line":
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			fmt.Println("*******")
			fmt.Println(line)
		}
	case "rune":
		for {
			r, _, err := reader.ReadRune()
			if err == io.EOF {
				break
			}
			fmt.Println("*******")
			fmt.Println(string(r))
		}
	case "string":
		for {
			var delim byte = '\n'
			str, err := reader.ReadString(delim)
			if err == io.EOF {
				break
			}
			fmt.Println("*******")
			fmt.Println(str)
		}
	case "byte":
		for {
			str, err := reader.ReadByte()
			if err == io.EOF {
				break
			}
			fmt.Println("*******")
			fmt.Println(string(str))
		}
	case "bytes":
		for {
			var delim byte = '\n'
			str, err := reader.ReadBytes(delim)
			if err == io.EOF {
				break
			}
			fmt.Println("*******")
			fmt.Println(string(str))
		}
	}
	fmt.Println("**************END****************")
	// fmt.Fprintf(os.Stdout, "%s\n", res)
}
