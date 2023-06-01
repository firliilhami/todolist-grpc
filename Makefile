gen:
	protoc --proto_path=proto --go_opt=module=todolist.firliilhami.com --go_out=. --go-grpc_opt=module=todolist.firliilhami.com --go-grpc_out=. proto/*.proto

clean:
	rm proto/*.go

server:
	go run cmd/server/main.go -port 1111

client:
	go run cmd/client/main.go -address 0.0.0.0:1111