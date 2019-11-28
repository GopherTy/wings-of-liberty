package main

import (
	"Go_Overwall/conf"
	"Go_Overwall/encryption"
)

func main() {
	cfg := conf.GetCfg()
	// test encode and decode
	b := []byte{30, 25, 7, 8, 2}
	encryArray := encryption.RandEncryArray()
	cipher := encryption.NewCipher(encryArray)
	cipher.Encode(b)
	sugar := cfg.Logger.Sugar()
	defer sugar.Sync()
	sugar.Info(b)
	cipher.Decode(b)
	sugar.Info(b)
}
