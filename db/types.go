package db

type Database interface {
	Create(string, interface{}) error
	Read(string, map[string]interface{}) (interface{}, error)
	ReadOne(string, map[string]interface{}) (interface{}, error)
	Update(string, map[string]interface{}, map[string]interface{}) error
}

type DBOperation struct {
	DB Database
}

const (
	MONGODB = "mongodb"
	MYSQL   = "mysql"
)

type MongoDetails struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	IP           string `json:"ip"`
	DatabaseName string `json:"dbname"`
}

type MySqlDetails struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	IP           string `json:"ip"`
	DatabaseName string `json:"dbname"`
}
