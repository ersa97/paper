package main

import (
	"log"
	"os"

	paper "github.com/ersa97/paper-test/controllers"
	"github.com/ersa97/paper-test/database"
	"github.com/ersa97/paper-test/routes"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	APPLICATION_PORT string
	DATABASE_URL     string
	db               *gorm.DB
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	APPLICATION_PORT = os.Getenv("APPLICATION_PORT")
	DATABASE_URL = os.Getenv("DATABASE_URL")

	db := database.Connection() // init database connection
	defer db.Close()

	paperService := paper.PaperService{
		DB: db,
	}

	routes.Mux(paperService)

}
