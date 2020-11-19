package model

import (
	database "g.ghn.vn/scte-common/godal"
)

func Create(dbType database.IDatabase, tableName string, dataStruct interface{}) (interface{}, error) {
	return dbType.CreateWithStruct(tableName, dataStruct)
}

func Get(dbType database.IDatabase, tableName string, limit int, offset int) ([]map[string]interface{}, error) {
	return dbType.GetAllToMap(tableName, limit, offset)
}

func Update(dbType database.IDatabase, tableName string, newValue map[string]interface{}, whereCondition map[string]interface{}) (interface{}, error) {
	return dbType.Update(tableName, newValue, whereCondition)
}

func Delete(dbType database.IDatabase, tableName string, whereCondition map[string]interface{}) (interface{}, error) {
	return dbType.Delete(tableName, whereCondition)
}
