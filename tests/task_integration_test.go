package tests

import (
	"context"
	"testing"

	pb "todolist.firliilhami.com/proto"

	"google.golang.org/grpc"
)

func TestTaskServiceIntegration(t *testing.T) {
	// Connect to the gRPC server
	conn, err := grpc.Dial("0.0.0.0:1111", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect to the gRPC server: %v", err)
	}
	defer conn.Close()

	// Create the TaskService client
	client := pb.NewTaskServiceClient(conn)

	// Create a new task
	createReq := &pb.CreateTaskRequest{
		Title:       "Test Task",
		Description: "This is a test task",
	}
	createRes, err := client.CreateTask(context.Background(), createReq)
	if err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	taskID := createRes.Id

	// Retrieve the task
	readReq := &pb.ReadTaskRequest{
		Id: taskID,
	}
	_, err = client.ReadTask(context.Background(), readReq)
	if err != nil {
		t.Fatalf("Failed to retrieve task: %v", err)
	}

	// Update the task
	updateReq := &pb.UpdateTaskRequest{
		Id:          taskID,
		Title:       "Updated Test Task",
		Description: "This is an updated test task",
	}
	_, err = client.UpdateTask(context.Background(), updateReq)
	if err != nil {
		t.Fatalf("Failed to update task: %v", err)
	}

	// Delete the task
	deleteReq := &pb.DeleteTaskRequest{
		Id: taskID,
	}
	_, err = client.DeleteTask(context.Background(), deleteReq)
	if err != nil {
		t.Fatalf("Failed to delete task: %v", err)
	}

	// Verify the task is deleted
	_, err = client.ReadTask(context.Background(), readReq)
	if err == nil {
		t.Fatal("Expected an error while retrieving deleted task, but got nil")
	}
}
