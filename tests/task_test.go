package tests

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"todolist.firliilhami.com/internal/models"
	"todolist.firliilhami.com/internal/service"
	pb "todolist.firliilhami.com/proto"
)

const expectedId uint32 = 1

func TestCreateTask(t *testing.T) {
	db := createMockDB()

	// initialize the TaskServer
	taskServer := service.NewTaskServer(db)

	// CreateTaskRequest
	req := &pb.CreateTaskRequest{
		Title:       "Test Title",
		Description: "Test Description",
	}

	// call the CreateTask Method
	res, err := taskServer.CreateTask(context.Background(), req)

	// assert that no error
	assert.Nil(t, err)

	// assert that the ID is 1
	assert.Equal(t, 1, int(res.Id))
}

func TestReadTask(t *testing.T) {
	db := createMockDB()

	// initialize the TaskServer
	taskServer := service.NewTaskServer(db)

	// create a task
	task := &models.Task{
		Title:       "Test Title",
		Description: "Test Description",
	}

	db.Create(task)

	// ReadTaskRequest
	req := &pb.ReadTaskRequest{
		Id: expectedId,
	}

	// call the ReadTask Method
	res, err := taskServer.ReadTask(context.Background(), req)

	// assert that no error
	assert.Nil(t, err)

	// assert that the ID is 1
	assert.Equal(t, 1, int(res.Id))

	// assert that the title is "Test Title"
	assert.Equal(t, "Test Title", res.Title)

	// assert that the description is "Test Description"
	assert.Equal(t, "Test Description", res.Description)
}

func TestUpdateTask(t *testing.T) {
	db := createMockDB()

	// initialize the TaskServer
	taskServer := service.NewTaskServer(db)

	// create a task
	task := &models.Task{
		Title:       "Test Title",
		Description: "Test Description",
	}

	db.Create(task)

	// CreateTaskRequest
	req := &pb.UpdateTaskRequest{
		Id:          expectedId,
		Title:       "updated title",
		Description: "updated description",
	}

	// call the UpdateTask Method
	res, err := taskServer.UpdateTask(context.Background(), req)

	// assert that no error
	assert.Nil(t, err)

	// assert that the ID is 1
	assert.Equal(t, 1, int(res.Id))

	// ReadTaskRequest
	readReq := &pb.ReadTaskRequest{
		Id: expectedId,
	}

	// call the ReadTask Method
	readRes, err := taskServer.ReadTask(context.Background(), readReq)

	// assert that no error
	assert.Nil(t, err)

	// assert that the ID is 1
	assert.Equal(t, 1, int(res.Id))

	// assert that the title is "updated title"
	assert.Equal(t, "updated title", readRes.Title)

	// assert that the description is "updated description"
	assert.Equal(t, "updated description", readRes.Description)

}

func TestDeleteTask(t *testing.T) {
	db := createMockDB()

	// initialize the TaskServer
	taskServer := service.NewTaskServer(db)

	// create a task
	task := &models.Task{
		Title:       "Test Title",
		Description: "Test Description",
	}

	db.Create(task)

	// DeleteTaskRequest
	req := &pb.DeleteTaskRequest{
		Id: 1,
	}

	// call the DeleteTask Method
	res, err := taskServer.DeleteTask(context.Background(), req)

	// assert that no error
	assert.Nil(t, err)

	// assert that the ID is not empty
	assert.Equal(t, true, res.Success)
}

func createMockDB() *gorm.DB {
	// Create an in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Automigrate the Task struct
	db.AutoMigrate(&models.Task{})

	return db
}
