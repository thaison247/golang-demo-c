package main

import (
	"context"
	"io"
	"log"
	"main/apis/employeepb"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())

	if err != nil {
		log.Fatal("error when dial to grpc server: ", err)
	}

	defer cc.Close()

	client := employeepb.NewEmployeeServiceClient(cc)

	callGetListEmployees(client)

}

func callGetEmployee(c employeepb.EmployeeServiceClient) {
	log.Println("Client calls -Get employee api")

	res, err := c.GetEmployee(context.Background(), &employeepb.EmployeeRequest{EmployeeId: 109})

	if err != nil {
		log.Fatalf("error when call get employee api %v", err)
	}

	log.Printf("Get employee api response: %v", res)
}

func callGetListEmployees(c employeepb.EmployeeServiceClient) {
	log.Println("Client calls -Get list employees api")

	stream, err := c.GetListEmployees(context.Background(), &employeepb.ListEmployeesRequest{Limit: 15, Offset: 0})

	if err != nil {
		log.Fatalf("Error when client call get list employees api: %v", err)
	}

	for {
		emp, recvErr := stream.Recv()

		if recvErr == io.EOF {
			log.Println("Server finished sending")
			return
		}

		log.Printf("Emloyee data: %v\n", emp)
	}
}
