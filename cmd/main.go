package main

import (
	"log"
	"os"

	"github.com/yurixtugal/Events/infrastructure/handler/response"
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = validateEnvironment()

	if err != nil {
		log.Fatal(err)
	}

	e := newHTTP(response.HTTPErrorHandler)

	dbPool, err := newDBConnection()

	if err != nil {
		log.Fatal(err)
	}

	_ = dbPool

	e.Start(":" + os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Fatal(err)
	}

}
