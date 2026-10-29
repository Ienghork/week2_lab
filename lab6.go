
package main

import "fmt"

func main() {
    var text string
    var key byte

    fmt.Print("Enter plaintext (one word): ")
    fmt.Scan(&text)

    fmt.Print("Enter key (one character): ")
    fmt.Scanf(" %c", &key)

    // Encrypt
    encrypted := ""

    for i := 0; i < len(text); i++ {
        encrypted += string(text[i] ^ key)
    }

    fmt.Println("Encrypted bytes:", []byte(encrypted))

    // Decrypt
    decrypted := ""

    for i := 0; i < len(encrypted); i++ {
        decrypted += string(encrypted[i] ^ key)
    }

    fmt.Println("Decrypted text:", decrypted)
}
