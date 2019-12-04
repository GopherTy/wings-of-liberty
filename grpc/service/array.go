package service

import (
	"context"
	"wings-of-liberty/encryption"
	array "wings-of-liberty/grpc/code"
)

const (
	// Address grpc server port
	Address = ":10000"
)

// Server get array server
type Server struct {
	Array *encryption.EncrypArray
}

// GetEncryptionArray grpc service,aim to send encryption array.
func (s *Server) GetEncryptionArray(ctx context.Context, req *array.ArrayRequest) (resp *array.ArrayResponse, err error) {
	resp = &array.ArrayResponse{}
	resp.Array = make([]byte, 0)
	for i := range s.Array {
		resp.Array = append(resp.Array, s.Array[i])
	}
	return
}
