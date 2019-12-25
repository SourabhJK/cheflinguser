package main

import(
	"github.com/cheflinguser/config"
	"github.com/cheflinguser/router"
	"github.com/cheflinguser/utilities"
	"github.com/cheflinguser/db"
)

func main(){
	config.ReadConfig()
	utilities.GetLogger()
	db.InjectDB()
	router.Route()

}