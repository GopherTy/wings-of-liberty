package main

import (
	"wings-of-liberty/config"
	"wings-of-liberty/encryption"
)

func main() {
	config := config.GetConfig()
	// test encode and decode
	b := []byte{30, 25, 7, 8, 2}
	encryArray := encryption.RandEncryArray()
	cipher := encryption.NewCipher(encryArray)
	cipher.Encrypt(b)
	sugar := config.Logger.Sugar()
	defer sugar.Sync()
	sugar.Info(b)
	cipher.Decrypt(b)
	sugar.Info(config.Freedom.Port)
}
