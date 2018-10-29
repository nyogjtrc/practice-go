
protoc:
	protoc -I message/ message/*.proto --go_out=plugins=grpc:message

build:
	CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o message-server ./server/main.go
	CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o message-client ./client/main.go

docker-build:
	docker build -t message-server .

clean:
	rm -rf message-server message-client
