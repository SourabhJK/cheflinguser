package db

import (
	"github.com/cheflinguser/config"
)

var DBOpt DBOperation

func InjectDB() {
	conf := config.GetConfig()

	switch conf.DbData.Dbtype {
	case MONGODB:
		mongoDetails = MongoDetails{
			Username:     conf.DbData.Username,
			Password:     conf.DbData.Password,
			IP:           conf.DbData.IP,
			DatabaseName: conf.DbData.DbName,
		}
		DBOpt.DB = mongoDetails
	case MYSQL:
		mysqlDetails = MySqlDetails{
			Username:     conf.DbData.Username,
			Password:     conf.DbData.Password,
			IP:           conf.DbData.IP,
			DatabaseName: conf.DbData.DbName,
		}
		DBOpt.DB = mysqlDetails
		EstablishMysqlSession()
	}
}
