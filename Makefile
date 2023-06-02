gen:
	protoc --proto_path=proto --go_opt=module=todolist.firliilhami.com --go_out=. --go-grpc_opt=module=todolist.firliilhami.com --go-grpc_out=. proto/*.proto

clean:
	rm proto/*.go

server:
	go run cmd/server/main.go
