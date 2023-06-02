package service

import (
	"context"
	"errors"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"todolist.firliilhami.com/internal/models"
	pb "todolist.firliilhami.com/proto"
)

// TaskServer is the server that provides Task service
type TaskServer struct {
	db *gorm.DB
	pb.UnimplementedTaskServiceServer
}

// NewTaskServer return a new TaskServer
func NewTaskServer(db *gorm.DB) *TaskServer {
	return &TaskServer{db: db}
}

// CreateTask is used for add new data
func (s *TaskServer) CreateTask(
	ctx context.Context,
	req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {

	// validate title
	if err := s.validateTitleRequest(req.Title); err != nil {
		log.Printf("Invalid argument for title: %v", err)
		return nil, err
	}

	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
	}

	res := s.db.Create(&task)

	if res.Error != nil {
		log.Printf("Failed to create task: %v", res.Error)
		return nil, res.Error
	}

	log.Printf("Created task with ID: %d", task.ID)
	response := &pb.CreateTaskResponse{
		Id: task.ID,
	}
	return response, nil
}

// get Task based on Task ID
func (s *TaskServer) ReadTask(ctx context.Context, req *pb.ReadTaskRequest) (*pb.ReadTaskResponse, error) {
	task := &models.Task{}
	// read from db based on id
	res := s.db.First(&task, req.Id)
	if res.Error != nil {
		log.Printf("Failed to read task: %v", res.Error)
		return nil, res.Error
	}

	response := &pb.ReadTaskResponse{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   timestamppb.New(task.CreatedAt),
	}

	log.Printf("read task ID: %d", task.ID)

	return response, nil
}

// updating task title or description
func (s *TaskServer) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	task := &models.Task{}

	res := s.db.First(task, req.Id)

	if res.Error != nil {
		log.Printf("Failed to read task: %v", res.Error)
		return nil, res.Error
	}

	// validate title
	if err := s.validateTitleRequest(req.Title); err != nil {
		log.Printf("Invalid argument for title: %v", err)
		return nil, err
	}

	task.Title = req.Title
	task.Description = req.Description

	res = s.db.Save(task)
	if res.Error != nil {
		log.Printf("Failed to update task: %v", res.Error)
		return nil, res.Error
	}

	log.Printf("update task ID: %d ", task.ID)

	response := &pb.UpdateTaskResponse{
		Id: task.ID,
	}

	return response, nil
}

// delete task by task Id
func (s *TaskServer) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	res := s.db.Delete(&models.Task{}, req.Id)

	if res.RowsAffected == 0 {
		err := errors.New("id is not found")
		log.Printf("Failed to delete task: %v", err)
		return nil, err
	} else if res.Error != nil {
		log.Printf("Failed to delete task: %v", res.Error)
		return nil, res.Error
	}
	log.Printf("the task is deleted ID: %d", req.Id)

	response := &pb.DeleteTaskResponse{
		Success: true,
	}
	return response, nil
}

func (s *TaskServer) validateTitleRequest(title string) error {
	if title == "" {
		return errors.New("title cannot be empty string")
	}
	return nil
}
