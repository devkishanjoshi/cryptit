package main

import (
	"fmt"

	"github.com/devkishanjoshi/cryptit/decrypt"
	"github.com/devkishanjoshi/cryptit/encrypt"
)

func main() {
	encryptStr := encrypt.Nimbus("Hello")
	fmt.Printf("Encrypted String: %v\n", encryptStr)

	decryptStr := decrypt.Nimbus(encryptStr)
	fmt.Printf("Decrypted String: %v\n", decryptStr)

}
