package main

import (
	"context"
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

	callGetEmployee(client)

}

func callGetEmployee(c employeepb.EmployeeServiceClient) {
	log.Println("Client calls -Get employee api")

	res, err := c.GetEmployee(context.Background(), &employeepb.EmployeeRequest{EmployeeId: 107})

	if err != nil {
		log.Fatalf("error when call get employee api %v", err)
	}

	log.Printf("Get employee api response: %v", res)
}
