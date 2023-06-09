// gateway
package main

import (
	"context"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc-test/pkg/pb"
)

// Person struct for use in my project
type Person struct {
	Name       string     `json:"name"`
	Age        int32      `json:"age"`
	PhoneMain  Phone      `json:"phone_main"`
	PhoneOther []Phone    `json:"phone_other"`
	Birthday   *time.Time `json:"birthday"`
}

type Phone struct {
	Name string                        `json:"name"`
	Type pb.GetPersonRequest_PhoneType `json:"type"`
}

func main() {
	e := echo.New()
	//create client for connect to grpc
	client := NewClient()
	e.GET("/person", client.GetPerson)

	e.Start(":1325")
}

type UserClient struct {
	client pb.PersonManagementClient
}

func NewClient() UserClient {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	//defer conn.Close()
	return UserClient{client: pb.NewPersonManagementClient(conn)}
}

func (user UserClient) GetPerson(c echo.Context) error {
	var p Person
	//string and int can normally send value
	p.Age = 30
	p.Name = "spider man"

	//when use grpc struct have to convert by use its function
	p.PhoneMain = Phone{Name: "0", Type: pb.GetPersonRequest_MOBILE}
	phoneMain := pb.GetPersonRequest_PhoneName{
		Name: p.PhoneMain.Name,
		Type: pb.GetPersonRequest_PhoneType(p.PhoneMain.Type),
	}

	//array
	p.PhoneOther = []Phone{{"1", pb.GetPersonRequest_HOME}, {"2", pb.GetPersonRequest_WORK}}
	phoneOthers := []*pb.GetPersonRequest_PhoneName{
		{
			Name: p.PhoneOther[0].Name,
			Type: p.PhoneOther[0].Type,
		},
		{
			Name: p.PhoneOther[1].Name,
			Type: p.PhoneOther[1].Type,
		},
	}

	//have to set time.Time type to timestamp of proto by using library
	//go get google.golang.org/protobuf/types/known/timestamppb
	t := time.Date(2000, 12, 12, 0, 0, 0, 0, time.Local)
	p.Birthday = &t
	timestampProto := timestamppb.New(t)

	req := &pb.GetPersonRequest{
		Name:       p.Name,
		Age:        p.Age,
		PhoneMain:  &phoneMain,
		PhoneOther: phoneOthers,
		Birthday:   timestampProto,
	}

	res, err := user.client.GetPerson(context.Background(), req)
	if err != nil {
		log.Fatalf("Add failed: %v", err)
	}
	log.Printf("Add result: %v", res)
	return c.JSON(200, res)
}
