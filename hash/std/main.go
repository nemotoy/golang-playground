package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	const (
		input1 = "The tunneling gopher digs downwards, "
		input2 = "unaware of what he will find."
	)

	first := sha256.New()
	first.Write([]byte(input1))
	fmt.Printf("%x\n", first.Sum(nil))

	marshaler, ok := first.(encoding.BinaryMarshaler)
	if !ok {
		log.Fatal("first does not implement encoding.BinaryMarshaler")
	}
	state, err := marshaler.MarshalBinary()
	if err != nil {
		log.Fatal("unable to marshal hash:", err)
	}

	second := sha256.New()

	unmarshaler, ok := second.(encoding.BinaryUnmarshaler)
	if !ok {
		log.Fatal("second does not implement encoding.BinaryUnmarshaler")
	}
	if err := unmarshaler.UnmarshalBinary(state); err != nil {
		log.Fatal("unable to unmarshal hash:", err)
	}

	first.Write([]byte(input2))

	fmt.Printf("%x\n", first.Sum(nil))
	fmt.Println(bytes.Equal(first.Sum(nil), second.Sum(nil)))

	// HMAC
	hash := hmac.New(sha512.New, []byte("myoriginkey"))
	id := string([]rune(hex.EncodeToString(hash.Sum(nil))))
	fmt.Printf("%s\n", id)
	id1 := string([]rune(hex.EncodeToString(hash.Sum(nil)))[0:12])
	fmt.Printf("%s\n", id1)
	hash.Write([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
	id2 := string([]rune(hex.EncodeToString(hash.Sum(nil)))[0:12])
	fmt.Printf("%s\n", id2)
}
