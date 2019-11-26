package main

import (
	"Go_Overwall/encryption"
	"fmt"
)

func main() {
	// test encode and decode
	b := []byte{30, 25, 7, 8, 2}
	encryArray := encryption.RandEncryArray()
	cipher := encryption.NewCipher(encryArray)
	cipher.Encode(b)
	fmt.Println(b)
	cipher.Decode(b)
	fmt.Println(b)
}
