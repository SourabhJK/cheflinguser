package main

import (
	userm "github.com/cheflinguser/chefling/usermanagement"
	"github.com/cheflinguser/config"
	"github.com/cheflinguser/db"
	"github.com/cheflinguser/router"
	"github.com/cheflinguser/utilities"
)

func main() {
	config.ReadConfig()
	utilities.GetLogger()
	db.InjectDB()

	conf := config.GetConfig()
	if conf.DbData.Dbtype == db.MYSQL {
		userm.UpdateUserModels()
	}

	router.Route()

}
