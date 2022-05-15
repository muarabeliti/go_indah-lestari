package main

import (
	// "myproject/config"
	"myproject/config"
	"myproject/routes"

	"github.com/go-playground/validator/v10"
)

func init() {
	// config.InitialMigration()
}

// var Validator *validator.Validate

func main() {

	config.Validator = validator.New()

	// create a new echo instance
	e := routes.New()

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
