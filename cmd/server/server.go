package server

import (
	"e-ticket/cmd/cli"
	"e-ticket/internal/config"
	"e-ticket/internal/router"
	appdatabase "e-ticket/pkg/database"
	appenviroment "e-ticket/pkg/enviroment"

	"log"
)

func init() {
	//$ go run main.go -migration=status
	//$ go run main.go -migration=up
	//$ go run main.go -migration=down
}

func Run(enviromenType appenviroment.EnviromentType) {

	//* Setup Enviroment
	appenviroment.Set(enviromenType)

	//* Load Config
	appConfig, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	//* Initialize Database
	db, err := appdatabase.NewPostgresDb(appConfig.Db.DbConnedtionString())
	if err != nil {
		log.Fatal(err)
	}

	//* DB Checking
	err = appdatabase.Check(db)
	if err != nil {
		log.Fatal(err)
	} else {
		defer appdatabase.Close(db)
	}
	var dbEntity appdatabase.DbEntity = appdatabase.DbEntity{PQ: db}

	//* Migration CLI
	var isApplied bool = cli.DbMigration(&dbEntity)
	if isApplied {
		return
	}

	//* Setup Route
	router := router.SetupRouter(&dbEntity)

	//* Server Running
	err = router.Run(appConfig.Server.IP + ":" + appConfig.Server.Port)
	if err != nil {
		log.Fatal(err)
	}

}
