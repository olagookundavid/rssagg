package main

/*
questions for Gerald:

parameters in HandlerCreateUser what of optional ones,
how does Go deal with null,
is this methods the way to go, orm and all,
how to host,
login dilema,
wait group in scraper file how does it work, does it mean that we use it to block because we don't get a chan
*/

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/olagookundavid/rssagg/handlers"
	"github.com/olagookundavid/rssagg/internal/database"
	"github.com/olagookundavid/rssagg/routes"
	"github.com/olagookundavid/rssagg/scraper"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT env variable missing")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL env variable missing")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database")
	}
	queries := database.New(conn)

	apiCfg := handlers.ApiConfig{
		DB: queries,
	}
	go scraper.StartScrapping(queries, 10, time.Minute)
	router := routes.CreateRouter(&apiCfg)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	// Start serving an http server
	fmt.Printf("Server starting on port %v\n", port)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}

/*
schema are migrations, run with 'goose postgres postgres://postgres:postgress@localhost:5432/rssagg up' for up migration and vice versa, cd to migration folder tho

queries are the sql changed to go code, run with 'sqlc generate'

handled null in posts.description check model.go and Scraper.go

*/
