gen:
	 protoc --go_out=pb --go-grpc_out=pb  --proto_path=proto proto/*.proto

clean:
	rm -rf pb/*

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

test:
	go test -cover -race ./...