package model

import (
	database "g.ghn.vn/scte-common/godal"
)

var (
	EMP_DEP string = "emp_dep"
)

func AddEmployeeToDeparmentWithMap(dbType database.IDatabase, mapData map[string]interface{}) (interface{}, error) {
	return dbType.Create(EMP_DEP, mapData)
}
