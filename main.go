package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/evidetta/db_migrations/config"
	"github.com/evidetta/db_migrations/models"
)

func main() {

	log.Println("Verifying configurations...")
	db_conf, err := config.NewDatabaseConfig(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL_MODE"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Configuration verified.")
	log.Println("Attempting to connect to database...")

	db, err := sql.Open("postgres", db_conf.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Verifying database connection...")
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database successfully.")

	user, err := models.CreateUser(db, "Elia", "London", time.Date(1988, 9, 14, 0, 0, 0, 0, time.UTC))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
}
