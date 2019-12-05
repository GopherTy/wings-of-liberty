package grpc

import (
	"log"
	"net"
	"wings-of-liberty/encryption"
	array "wings-of-liberty/grpc/code"
	"wings-of-liberty/grpc/service"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

// Run run grpc service
func Run() {
	listener, err := net.Listen("tcp", service.Address)
	if err != nil {
		log.Fatalln("listen grpc server  fail ", err, service.Address)
	}

	// regist array server
	s := grpc.NewServer()
	arr := encryption.RandEncryArray()
	array.RegisterArrayServer(s, &service.Server{
		Array: arr,
	})
	reflection.Register(s)

	log.Println("GRPC service running ...")

	if err := s.Serve(listener); err != nil {
		log.Fatalf("grpc server start fail ,%v", err)
	}
}
