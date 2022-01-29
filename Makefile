proto:
	protoc -I pkg/helloworld --go_out=plugins=grpc:pkg/helloworld pkg/helloworld/helloworld.proto
docker-build:
	docker build . -t index.alauda.cn/alaudak8s/asm-test-image
docker-push:
	docker push index.alauda.cn/alaudak8s/asm-test-image
build:
	rm -rf bin
	go build -o bin/grpc-client cmd/grpc-client/main.go
	go build -o bin/grpc-server cmd/grpc-server/main.go
	go build -o bin/http-server cmd/http-server/main.go
	go build -o bin/h2c-server cmd/h2c-server/main.go
