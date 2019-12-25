package config

type Config struct{
	DbData Db `json:"dbdata"`
	TestToken string `json:"testkey"`
}

type Db struct{
	Dbtype string `json:"db"`
	Username string `json:"username"`
	Password string `json:"password"`
	IP string `json:"ip"`
	DbName string `json:"dbname"`
}