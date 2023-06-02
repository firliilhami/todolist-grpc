package main

import (
	"fmt"
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
	port := 1111

	log.Printf("start server on port %d", port)

	// dsn database

	dsn := "user=postgres password=postgres host=db port=5432 dbname=postgres sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("DB connected")

	// Auto-migrate the table
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate table: %v", err)
	}

	taskServer := service.NewTaskServer(db)
	grpcServer := grpc.NewServer()

	pb.RegisterTaskServiceServer(grpcServer, taskServer)

	address := fmt.Sprintf("0.0.0.0:%d", port)

	// Enable the reflection API (it is used for grpcurl)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
