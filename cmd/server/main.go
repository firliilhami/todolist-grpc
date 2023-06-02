package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"todolist.firliilhami.com/internal/wire"
	pb "todolist.firliilhami.com/proto"
)

func main() {

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
		log.Fatal("cannot start server: ", err)
	}

	log.Println("the server is running on: " + address)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}

func newDatabaseURL() string {
	env := os.Getenv("env")
	var host string
	var port int32

	if env == "" {
		host = "localhost"
		port = 2222
	} else {
		host = "db"
		port = 5432
	}

	return fmt.Sprintf("host=%s port=%d user=postgres password=postgres dbname=postgres sslmode=disable", host, port)
}
func newDBConn(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	return db, nil
}
