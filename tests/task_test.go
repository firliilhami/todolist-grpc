package tests

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"todolist.firliilhami.com/internal/models"
	"todolist.firliilhami.com/internal/service"
	pb "todolist.firliilhami.com/proto"
)

func TestCreateTask(t *testing.T) {
	db := createMockDB()

	// initialize the TaskServer
	taskServer := service.NewTaskServer(db)
	type TestCase struct {
		name        string
		title       string
		description string
		code        codes.Code
	}

	testCases := []TestCase{
		{
			name:        "valid_arguments",
			title:       "valid title",
			description: "valid description",
			code:        codes.OK,
		},
		{
			name:        "invalid_arguments_empty_title",
			title:       "",
			description: "empty string for title",
			code:        codes.InvalidArgument,
		},
	}

	var tc TestCase
	var req *pb.CreateTaskRequest
	var res *pb.CreateTaskResponse
	var err error

	for i := range testCases {
		tc = testCases[i]

		req = &pb.CreateTaskRequest{
			Title:       tc.title,
			Description: tc.description,
		}

		res, err = taskServer.CreateTask(context.Background(), req)

		if tc.code == codes.OK {
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.NotEmpty(t, res.Id)
		} else {
			assert.Error(t, err)
			assert.Nil(t, res)
		}

	}
}

func TestReadTask(t *testing.T) {
	db := createMockDB()

	// initialize the TaskServer
	taskServer := service.NewTaskServer(db)
	type TestCase struct {
		name string
		id   uint32
		code codes.Code
	}

	testCases := []TestCase{
		{
			name: "valid_arguments",
			id:   1,
			code: codes.OK,
		},
		{
			name: "invalid_arguments_no_record_id",
			id:   1000,
			code: codes.InvalidArgument,
		},
	}

	// create task request to save data in database
	createReq := &pb.CreateTaskRequest{
		Title:       "test title",
		Description: "test description",
	}

	_, _ = taskServer.CreateTask(context.Background(), createReq)

	var tc TestCase
	var req *pb.ReadTaskRequest
	var res *pb.ReadTaskResponse
	var err error

	for i := range testCases {
		tc = testCases[i]

		req = &pb.ReadTaskRequest{
			Id: tc.id,
		}

		res, err = taskServer.ReadTask(context.Background(), req)

		if tc.code == codes.OK {
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.NotEmpty(t, res.Id)
			assert.Equal(t, createReq.Title, res.Title)
			assert.Equal(t, createReq.Description, res.Description)
		} else {
			assert.Error(t, err)
			assert.Nil(t, res)
		}

	}
}

func TestUpdateTask(t *testing.T) {
	db := createMockDB()

	// initialize the TaskServer
	taskServer := service.NewTaskServer(db)
	type TestCase struct {
		name        string
		id          uint32
		Title       string
		Description string
		code        codes.Code
	}

	testCases := []TestCase{
		{
			name:        "valid_arguments",
			id:          1,
			Title:       "title is updated",
			Description: "description is updated",
			code:        codes.OK,
		},
		{
			name:        "invalid_arguments_empty_title",
			id:          1,
			Title:       "",
			Description: "updated with empty title",
			code:        codes.InvalidArgument,
		},
		{
			name:        "invalid_arguments_no_record_id",
			id:          10000,
			Title:       "title is updated",
			Description: "description is updated",
			code:        codes.InvalidArgument,
		},
	}

	// create task request to save data in database
	createReq := &pb.CreateTaskRequest{
		Title:       "test title",
		Description: "test description",
	}

	_, _ = taskServer.CreateTask(context.Background(), createReq)

	var tc TestCase
	var readRes *pb.ReadTaskResponse
	var req *pb.UpdateTaskRequest
	var res *pb.UpdateTaskResponse
	var err error

	for i := range testCases {
		tc = testCases[i]

		req = &pb.UpdateTaskRequest{
			Id:          tc.id,
			Title:       tc.Title,
			Description: tc.Description,
		}

		res, err = taskServer.UpdateTask(context.Background(), req)

		if tc.code == codes.OK {
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.NotEmpty(t, res.Id)
		} else {
			assert.Error(t, err)
			assert.Nil(t, res)
		}

		// read the update task
		readRes, _ = taskServer.ReadTask(context.Background(), &pb.ReadTaskRequest{Id: 1})

		if tc.code == codes.OK {
			assert.Equal(t, tc.Title, readRes.Title)
			assert.Equal(t, tc.Description, readRes.Description)
		}

	}
}

func TestDeleteTask(t *testing.T) {
	db := createMockDB()

	// initialize the TaskServer
	taskServer := service.NewTaskServer(db)
	type TestCase struct {
		name string
		id   uint32
		code codes.Code
	}

	testCases := []TestCase{
		{
			name: "valid_arguments",
			id:   1,
			code: codes.OK,
		},
		{
			name: "invalid_arguments_no_record_id",
			id:   10000,
			code: codes.InvalidArgument,
		},
	}

	// create task request to save data in database
	createReq := &pb.CreateTaskRequest{
		Title:       "test title",
		Description: "test description",
	}

	_, _ = taskServer.CreateTask(context.Background(), createReq)

	var tc TestCase
	var req *pb.DeleteTaskRequest
	var res *pb.DeleteTaskResponse
	var err error

	for i := range testCases {
		tc = testCases[i]

		req = &pb.DeleteTaskRequest{
			Id: tc.id,
		}

		res, err = taskServer.DeleteTask(context.Background(), req)

		if tc.code == codes.OK {
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, true, res.Success)
		} else {
			log.Println(err)
			assert.Error(t, err)
			assert.Nil(t, res)
		}

	}
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
