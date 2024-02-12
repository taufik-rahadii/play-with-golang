package main

import (
	"log"
	"taufikRahadi/fiber-boilerplate/configs"
	"taufikRahadi/fiber-boilerplate/database"
)

func main() {
	configs.InitEnv("./.env")
	database.Connect()
	api := configs.InitApi()

	log.Fatal(api.Listen(":3000"))
}
