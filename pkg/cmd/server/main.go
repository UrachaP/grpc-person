// server.go
// grpc
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc-test/pkg/pb"
)

type server struct {
	pb.UnsafePersonManagementServer
}

func (s *server) GetPerson(ctx context.Context, req *pb.GetPersonRequest) (*pb.GetPersonResponse, error) {
	birthdayString := req.Birthday.String()
	birthdayTime := req.Birthday.AsTime()
	result := fmt.Sprintf("Person name is %s, age is %d, phone is %v,%v and %v, birthday string is %s and birthday time is %v",
		req.Name,
		req.Age,
		req.PhoneMain.Name, req.PhoneMain.Type,
		req.PhoneOther,
		birthdayString,
		birthdayTime,
	)
	return &pb.GetPersonResponse{Status: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterPersonManagementServer(s, &server{})

	fmt.Println("Server listening on port 8080")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
