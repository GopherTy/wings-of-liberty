package main

import (
	"log"
	"net"
	array "wings-of-liberty/grpc/code"
	"wings-of-liberty/grpc/service"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	listener, err := net.Listen("tcp", service.Address)
	if err != nil {
		log.Fatalln("listen grpc server  fail ", err, service.Address)
	}
	// regist array server
	s := grpc.NewServer()
	array.RegisterArrayServer(s, &service.Server{})
	reflection.Register(s)
	log.Println("grpc array service running ...")
	if err := s.Serve(listener); err != nil {
		log.Fatalln("grpc server start fail ", err)
	}
}
