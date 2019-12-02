package main

import (
	"net"
	"strconv"
	"wings-of-liberty/config"
	"wings-of-liberty/encryption"
	"wings-of-liberty/local"
)

const (
	// DefaultListenAddr default local address
	DefaultListenAddr = ":7448"
)

func main() {
	cfg := config.GetConfig()
	sugar := cfg.Logger.Sugar()
	defer sugar.Sync()

	arr := encryption.RandEncryArray()

	laddrStr := cfg.Freedom.LocalAddr + ":" +
		strconv.Itoa(cfg.Freedom.LocalPort)
	raddrStr := cfg.Freedom.RemotoAddr + ":" +
		strconv.Itoa(cfg.Freedom.RemotoPort)

	laddr, err := net.ResolveTCPAddr("tcp", laddrStr)
	if err != nil {
		sugar.Fatal(err)
	}
	remote, err := net.ResolveTCPAddr("tcp", raddrStr)
	if err != nil {
		sugar.Fatal(err)
	}
	client := local.NewClient(arr, laddr, remote)
	sugar.Infof("client running, local address is %s, remoto address is %s",
		laddrStr,
		raddrStr,
	)

	err = client.Listen()
	if err != nil {
		return
	}
}
