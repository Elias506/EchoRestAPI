package main

import (
	"github.com/elias506/EchoRestAPI/server"
	"log"
)

const (
	fileDB string = "./repository/users.json"
)

func main() {
	app := server.NewApp()
	app.InitFileDB(fileDB)
	port := "8080"
	if err := app.Run(port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
