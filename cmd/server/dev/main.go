package main

import (
	"e-ticket/cmd/server"
	appenviroment "e-ticket/pkg/enviroment"
)

func main() {
	server.Run(appenviroment.Development)
}
