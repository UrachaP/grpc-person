grpc:
	protoc --go_out=./pkg/ --go-grpc_out=./pkg/ pkg/pb_gen/Person.proto

