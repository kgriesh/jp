package main

import (
	"fmt"
	"jp/app"
	"jp/app/db"
	"log"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

func main() {

	addr := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	host := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	database, err := db.NewDbService(host, dbUser, dbPassword, dbName, dbPort)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.GetConnection().Close()

	dinoService := app.NewDinoService(database)

	handler := app.NewHandler(dinoService, &logger)
	err = http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatalf("Could start app: %v", err)
	}

}
