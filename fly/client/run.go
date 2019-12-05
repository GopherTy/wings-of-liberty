package client

import (
	"context"
	"net"
	"strconv"
	"wings-of-liberty/config"
	"wings-of-liberty/encryption"
	array "wings-of-liberty/grpc/code"
	"wings-of-liberty/local"

	"google.golang.org/grpc"
)

// Run ...
func Run() {
	cfg := config.GetConfig()
	sugar := cfg.Logger.Sugar()
	defer sugar.Sync()

	arr, err := grpcArrayService(cfg)
	if err != nil {
		sugar.Fatal(err)
	}
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
	// create a client
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

// grpcArrayService grpc service get encryption array
func grpcArrayService(cfg *config.Config) (arr *encryption.EncrypArray, err error) {
	sugar := cfg.Logger.Sugar()
	defer sugar.Sync()

	addr := cfg.Freedom.RemotoAddr + ":" + "10000"

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		sugar.Fatal("grpc service start fail ", err)
	}
	gClient := array.NewArrayClient(conn)
	resp, err := gClient.GetEncryptionArray(context.Background(), &array.ArrayRequest{})
	if err != nil {
		sugar.Fatal("recive grpc service data fail ", err)
	}
	if resp.Array == nil {
		sugar.Fatal("array is nil")
	}
	arr = &encryption.EncrypArray{}
	for i, v := range resp.Array {
		arr[i] = v
	}
	return
}
