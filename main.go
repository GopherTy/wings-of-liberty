package main

import (
	"Go_Overwall/config"
	"Go_Overwall/encryption"
)

func main() {
	config := config.GetConfig()
	// test encode and decode
	b := []byte{30, 25, 7, 8, 2}
	encryArray := encryption.RandEncryArray()
	cipher := encryption.NewCipher(encryArray)
	cipher.Encode(b)
	sugar := config.Logger.Sugar()
	defer sugar.Sync()
	sugar.Info(b)
	cipher.Decode(b)
	sugar.Info(config.Freedom.Port)
}
