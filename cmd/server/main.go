package main

import (
	"flag"
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
	port := flag.Int("port", 0, "the server port")
	flag.Parse()

	log.Printf("start server on port %d", *port)

	// database
	dsn := "host=0.0.0.0 user=postgres password=postgres dbname=postgres port=1717 sslmode=disable"

	var err error
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

	address := fmt.Sprintf("0.0.0.0:%d", *port)

	// Enable the reflection API
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
