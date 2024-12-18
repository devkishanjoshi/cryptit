package main

import (
	"fmt"

	"github.com/devkishanjoshi/cryptit/encrypt"
)

func main() {
	encryptStr := encrypt.Nimbus("Hello World")
	fmt.Printf("Encrypted String: %v\n", encryptStr)
}
