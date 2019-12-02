package main

import (
	"log"
	"net"
	"strconv"
	"wings-of-liberty/config"
	"wings-of-liberty/encryption"
	"wings-of-liberty/remote"
)

func main() {
	cfg := config.GetConfig()
	sugar := cfg.Logger.Sugar()
	defer sugar.Sync()

	arr := encryption.RandEncryArray()
	laddrStr := ":" + strconv.Itoa(cfg.Freedom.RemotoPort)
	local, err := net.ResolveTCPAddr("tcp", laddrStr)
	if err != nil {
		log.Fatalln(err)
	}

	s := remote.NewServer(arr, local)
	sugar.Info("Listening ... ")
	err = s.Listen()
	if err != nil {
		log.Fatalln(err)
	}
}
