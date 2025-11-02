package cli

import (
	appdatabase "e-ticket/pkg/database"
	applogger "e-ticket/pkg/logger"

	"flag"
	"log"

	"github.com/pressly/goose/v3"
)

const (
	migrationPath string = "./migration"
)

func DbMigration(db *appdatabase.DbEntity) bool {
	var migrationFlag string
	flag.StringVar(&migrationFlag, "migration", "", "Available Command: up/down/status")

	flag.Parse()

	if migrationFlag == "" {
		return false
	}

	switch migrationFlag {
	case "up":
		err := goose.Up(db.PQ, migrationPath)
		if err != nil {
			log.Fatal(err)
		}
		return true

	case "down":
		err := goose.Down(db.PQ, migrationPath)
		if err != nil {
			log.Fatal(err)
		}
		return true

	case "status":
		err := goose.Status(db.PQ, migrationPath)
		if err != nil {
			log.Fatal(err)
		}
		return true

	default:
		applogger.Info("command not found")
		return true
	}
}
