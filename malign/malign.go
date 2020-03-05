package main

import (
	"fmt"
	"unsafe"
)

// nolint: structcheck, unused, maligned
type myStruct struct {
	myInt   bool
	myFloat float64
	myBool  int32
}

// nolint: structcheck, unused
type myStructOptimized struct {
	myFloat float64
	myInt   int32
	myBool  bool
}

func main() {
	fmt.Println(unsafe.Sizeof(myStruct{}))          // unordered 24 bytes
	fmt.Println(unsafe.Sizeof(myStructOptimized{})) // ordered 16 bytes
}
