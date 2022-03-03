package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	text := "Hello World - Go!"

	encodedText := base64.StdEncoding.EncodeToString([]byte(text));

	fmt.Printf("Encoded Text: %s\n", encodedText)

	decodedText, err := base64.StdEncoding.DecodeString(encodedText);
    if err != nil {
        panic(err)
    }

	fmt.Printf("Decoded Text: %s\n", decodedText)
}