package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"todolist.firliilhami.com/internal/wire"
	pb "todolist.firliilhami.com/proto"
)

func main() {
	// Initialize task server
	taskServer, err := wire.InitializeTaskServer()
	if err != nil {
		log.Fatal("cannot initialize server: ", err)
	}
	grpcServer := grpc.NewServer()

	// Enable the reflection API (so grpcurl can use the service)
	reflection.Register(grpcServer)

	pb.RegisterTaskServiceServer(grpcServer, taskServer)

	address := "0.0.0.0:1111"

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	log.Println("starting server on: " + address)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
