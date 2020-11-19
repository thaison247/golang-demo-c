package utils

var Global map[string]interface{} = make(map[string]interface{})
var AppConfig *HoconConfig

const POSTGRES_ENTITY = "POSTGRES_ENTITY"
