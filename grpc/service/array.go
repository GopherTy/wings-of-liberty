package service

import (
	"context"
	"fmt"
	"wings-of-liberty/encryption"
	array "wings-of-liberty/grpc/code"
)

const (
	// Address grpc server port
	Address = "127.0.0.1:10000"
)

// Server get array server
type Server struct {
}

// GetEncryptionArray grpc service,aim to send encryption array.
func (s *Server) GetEncryptionArray(ctx context.Context, req *array.ArrayRequest) (resp *array.ArrayResponse, err error) {
	// rand encryption array
	arr := encryption.RandEncryArray()
	resp = &array.ArrayResponse{}
	resp.Array = make([]byte, 0)
	for i := range arr {
		resp.Array = append(resp.Array, arr[i])
	}
	fmt.Println(resp.Array)
	return
}
