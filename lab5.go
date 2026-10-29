
package main

import (
	"fmt"
	"encoding/base64"
	"encoding/hex"
)

func main() {
	var text string

	fmt.Print(" Enter text: ")
	fmt.Scanln( &text)

	for i := 0; i< len(text); i++ {
		fmt.Printf("%08b ", text[i] )
	}
	fmt.Println()
	fmt.Println("Hexadecimal:", hex.EncodeToString([]byte(text)))

	fmt.Println("Base64: ", base64.StdEncoding.EncodeToString( []byte(text) ))
}