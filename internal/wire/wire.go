//go:build wireinject
// +build wireinject

package wire

import (
	"fmt"
	"log"
	"os"

	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"todolist.firliilhami.com/internal/models"
	"todolist.firliilhami.com/internal/service"
)

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

	// Auto-migrate the table
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate table: %v", err)
		return nil, err
	}

	return db, nil
}

func InitializeTaskServer() (*service.TaskServer, error) {
	wire.Build(newDBConn, newDatabaseURL, service.NewTaskServer)
	return nil, nil
}
