package main

import "fmt"

var (
	version string
	date    string
)

func main() {
	fmt.Println("app version: ", version)
	fmt.Println("build date: ", date)
}
