package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"todolist.firliilhami.com/internal/models"
	"todolist.firliilhami.com/internal/service"
	pb "todolist.firliilhami.com/proto"
)

func main() {
	// dsn database
	dsn := "user=postgres password=postgres host=db port=5432 dbname=postgres sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Auto-migrate the table
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate table: %v", err)
	}

	taskServer := service.NewTaskServer(db)
	grpcServer := grpc.NewServer()

	pb.RegisterTaskServiceServer(grpcServer, taskServer)

	address := "0.0.0.0:1111"

	// Enable the reflection API (it is used for grpcurl)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	log.Println("the server is running on: " + address)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
