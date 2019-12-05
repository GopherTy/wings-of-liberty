package server

import (
	"context"
	"net"
	"strconv"
	"wings-of-liberty/config"
	"wings-of-liberty/encryption"
	array "wings-of-liberty/grpc/code"
	"wings-of-liberty/remote"

	"google.golang.org/grpc"
)

// Run running app
func Run() {
	cfg := config.GetConfig()
	sugar := cfg.Logger.Sugar()
	defer sugar.Sync()

	// grpc service to get encryption array
	arr, err := grpcArrayService(cfg)
	if err != nil {
		sugar.Fatal(err)
	}

	laddrStr := ":" + strconv.Itoa(cfg.Freedom.RemotoPort)
	local, err := net.ResolveTCPAddr("tcp", laddrStr)
	if err != nil {
		sugar.Fatal(err)
	}

	s := remote.NewServer(arr, local)
	sugar.Info("Server running ... ")
	err = s.Listen()
	if err != nil {
		sugar.Fatal(err)
	}
}

// grpcArrayService grpc service get encryption array
func grpcArrayService(cfg *config.Config) (arr *encryption.EncrypArray, err error) {
	sugar := cfg.Logger.Sugar()
	defer sugar.Sync()

	conn, err := grpc.Dial(":10000", grpc.WithInsecure())
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
