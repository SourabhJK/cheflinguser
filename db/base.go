package db

import(
	"github.com/cheflinguser/config"
)


var DBOpt DBOperation


func InjectDB(){
	conf := config.GetConfig()

	switch conf.DbData.Dbtype{
	case MONGODB:
		var mongo = MongoDetails{
			Username: conf.DbData.Username,
			Password: conf.DbData.Password,
			IP: conf.DbData.IP,
			DatabaseName: conf.DbData.DbName,
		}
		DBOpt.DB = mongo
	// case MYSQL:
	// 	var mysql = MySqlDetails{
	// 		Username: conf.DbData.Username,
	// 		Password: conf.DbData.Password,
	// 		IP: conf.DbData.IP,
	// 		DatabaseName: conf.DbData.DbName,
	// 	}
	// 	DBOpt.DB = mysql

	}
}