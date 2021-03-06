package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/evidetta/web-api/config"
	"github.com/evidetta/web-api/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {

	log.Println("Verifying configurations...")
	db_conf, err := config.NewDatabaseConfig(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL_MODE"))
	if err != nil {
		log.Fatal(err)
	}

	api_conf, err := config.NewAPIConfig(os.Getenv("API_PORT"), os.Getenv("API_PAGE_SIZE"), os.Getenv("API_RUN_MIGRATIONS"))
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

	if api_conf.RunMigrations == true {
		log.Println("Running migrations...")

		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			log.Fatal(err)
		}

		m, err := migrate.NewWithDatabaseInstance("file:///migrations", "postgres", driver)
		if err != nil {
			log.Fatal(err)
		}

		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}

		log.Println("Migrations run successfully.")
	}

	log.Println("Setting up server...")
	handlers.Init(db, api_conf.PageSize)

	r := mux.NewRouter()
	r.HandleFunc("/ping", handlers.Ping).Methods("GET")

	r.HandleFunc("/users", handlers.GetUsers).Methods("GET").Queries("page", "{page}")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")

	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/user", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user", handlers.DeleteUser).Methods("DELETE")

	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	srv := &http.Server{
		Addr: api_conf.GetAPIAddress(),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	log.Println("Server running.")

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down...")
	os.Exit(0)
}
