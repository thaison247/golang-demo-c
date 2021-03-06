package server

import (
	"context"
	"encoding/json"
	"log"
	"main/apis/employeepb"
	"main/model"
	"main/utils"
	"time"

	database "g.ghn.vn/scte-common/godal"
)

type EmployeeServer struct{}

func (e *EmployeeServer) GetEmployee(ctx context.Context, req *employeepb.EmployeeRequest) (
	*employeepb.EmployeeResponse,
	error,
) {
	log.Println("Get employee service")
	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)

	rs, err := model.GetEmployeeByIdV2(dbType, int(req.GetEmployeeId()))
	if err != nil {
		return nil, err
	}

	row := rs[0]

	//convert map[string]interface{} to json
	jsonString, _ := json.Marshal(row)

	//convert json to struct
	emp := &employeepb.EmployeeResponse{}
	json.Unmarshal(jsonString, emp)

	return emp, nil
}

func (e *EmployeeServer) GetListEmployees(req *employeepb.ListEmployeesRequest,
	stream employeepb.EmployeeService_GetListEmployeesServer) error {
	limit := req.GetLimit()
	offset := req.GetOffset()

	dbtype := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	rs, err := model.GetEmployees(dbtype, int(limit), int(offset))

	if err != nil {
		return err
	}

	for _, val := range rs {
		//convert map[string]interface{} to json
		jsonString, _ := json.Marshal(val)
		//convert json to struct
		emp := &employeepb.EmployeeResponse{}
		json.Unmarshal(jsonString, emp)

		if err := stream.Send(emp); err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

func (e *EmployeeServer) CreateEmployee(ctx context.Context, req *employeepb.EmployeeRequest) (*employeepb.CreateEmployeeResponse, error) {
	employee := make(map[string]interface{})
	jsonString, _ := json.Marshal(req)
	json.Unmarshal(jsonString, &employee)

	delete(employee, "employee_id")

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	_, err := model.CreateWithMap(dbType, employee)

	if err != nil {
		log.Fatalf("error when create employee: %v\n", err)
		resp := &employeepb.CreateEmployeeResponse{StatusMsg: "Fail"}
		return resp, err
	}

	resp := &employeepb.CreateEmployeeResponse{StatusCode: 200, StatusMsg: "Success"}

	return resp, nil
}

func (e *EmployeeServer) DeleteEmployee(ctx context.Context, req *employeepb.EmployeeRequest) (*employeepb.DeleteEmployeeResponse, error) {
	employee := make(map[string]interface{})
	jsonString, _ := json.Marshal(req)
	json.Unmarshal(jsonString, &employee)

	log.Println(employee)

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	_, err := model.Delete(dbType, "employees", employee)

	if err != nil {
		log.Fatalf("Error when delete employee: %v\n", err)
		resp := &employeepb.DeleteEmployeeResponse{
			StatusMsg: "Fail",
		}
		return resp, err
	}

	resp := &employeepb.DeleteEmployeeResponse{
		StatusCode: 200,
		StatusMsg:  "Success",
	}

	return resp, nil
}

func (e *EmployeeServer) UpdateEmployee(ctx context.Context, req *employeepb.EmployeeRequest) (*employeepb.EmployeeResponse, error) {
	reqData := make(map[string]interface{})
	jsonString, _ := json.Marshal(req)
	json.Unmarshal(jsonString, &reqData)

	dbtype := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	_, err := model.Update(dbtype, "employees", reqData, map[string]interface{}{"employee_id": reqData["employee_id"]})

	if err != nil {
		return nil, err
	}

	employeeId := int(reqData["employee_id"].(float64))

	log.Println(employeeId)

	res, err := model.Get(dbtype, "employees", 1, 0)

	if err != nil {
		return nil, err
	}

	log.Println(res[0])

	// row := res[0]

	// log.Println(row)

	//convert map[string]interface{} to json
	jsonString, _ = json.Marshal(res[0])

	//convert json to struct
	emp := &employeepb.EmployeeResponse{}
	json.Unmarshal(jsonString, emp)

	return emp, nil
}
