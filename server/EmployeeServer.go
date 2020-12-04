package server

import (
	"context"
	"log"
	"main/apis/employeepb"
	"main/model"
	"main/utils"

	database "g.ghn.vn/scte-common/godal"
)

type EmployeeServer struct{}

func (e *EmployeeServer) GetEmployee(ctx context.Context, req *employeepb.EmployeeRequest) (*employeepb.EmployeeResponse, error) {
	log.Println("Get employee service")
	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	_, err := model.GetEmployeeById(dbType, int(req.GetEmployeeId()))
	if err != nil {
		return nil, err
	}

	emp := &employeepb.EmployeeResponse{
		FullName: "full name",
	}

	return emp, nil
}
