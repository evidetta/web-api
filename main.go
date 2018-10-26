package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/evidetta/db_migrations/config"
)

func main() {

	log.Println("Verifying configurations...")
	db_conf, err := config.NewDatabaseConfig(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL_MODE"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Configuration verified.")
	log.Println("Attempting to connect to database...")

	//127.0.0.1 readwrite password db disable
	db, err := sql.Open("postgres", db_conf.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database successfully.")
}
