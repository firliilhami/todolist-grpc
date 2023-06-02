# Use the official Golang base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o ./server ./cmd/server

# Expose the gRPC server port
EXPOSE 1111

# Set the command to run the server executable
CMD ["./server"]