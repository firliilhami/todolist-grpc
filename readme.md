# Todolist app
This is a Todo app in golang using gRPC for server-client communication and implementing Postgresql for database, which allows users to create, read, update, and delete tasks

Features:
1. CreateTask: Users can create new tasks by providing a title and description. Each task is assigned a unique identifier for easy reference.
2. ReadTask: Users can retrieve task details by providing the task ID. The app returns the title, description, and other relevant information associated with the task.
3. UpdateTask: Users can update the title and description of an existing task by providing the task ID along with the updated information.
4. DeleteTask: Users can delete a task by specifying the task ID. This removes the task from the database.

# How to run gRPC server
1. Clone this github repository
2. Open your terminal and change the directory to this project folder
3. We run the server using docker-compose, ensure that you have Docker and Docker Compose installed on your system.
4. Run ```docker-compose up``` on your terminal, it will take some time when run for the first time.
5. If it run successully, it will log ```the server is running on: 0.0.0.0:1111```

# Test CRUD operations
Once the containers are running, you can use ```grpcurl``` to test the CRUD operations. Ensure that you have grpcurl installed on your local machine.

For example, to create a new task, you can use the following command:
```
grpcurl -plaintext -d '{"title": "Review a PR", "description": "Reviewing the PR for feature X"}' 0.0.0.0:1111 task_package.TaskService/CreateTask
```
```
expected response:
{
    "id": uint32,
}
```
To get a task by ID, use the following command:
```
grpcurl -plaintext -d '{"id": 1}' 0.0.0.0::1111 task_package.TaskService/ReadTask
```
```
expected response:
{
  "id": uint32,
  "title": string,
  "description": string,
  "createdAt": Timestamp
}
```
To update a task, use the following command:
```
grpcurl -plaintext -d '{"id": 1, "title": "Review a PR Updated", "description": "Reviewing the PR for feature X updated"}' 0.0.0.0::1111 task_package.TaskService/UpdateTask
```
```
expected response:
{
    "id": uint32,
}
```

to delete a task by ID, use the following command:
```
grpcurl -plaintext -d '{"id": 1}' 0.0.0.0::1111 task_package.TaskService/DeleteTask
```
```
expected response:
{
    "success": bool,
}
```


