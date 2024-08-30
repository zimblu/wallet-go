package main

import (
	"fmt"
	"log"

	"github.com/tyler-smith/go-bip39"
)

func create_mnemonic() {
	// generate Entropy, y=32*x 128<=y<=256
	b, err := bip39.NewEntropy(128)
	if err != nil {
		log.Panic("failed to NewEntropy:", err, b)
	}
	// generate Mnemonic
	nm, err := bip39.NewMnemonic(b)
	if err != nil {
		log.Panic("failed to NewMnemonic:", err)
	}
	fmt.Println(nm)
}

func main() {
	create_mnemonic()
}
