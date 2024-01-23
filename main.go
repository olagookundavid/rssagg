package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT env variable missing")
	}

	// dbUrl := os.Getenv("DB_URI")
	// if dbUrl == "" {
	// 	log.Fatal("DB_URL env variable missing")
	// }

	// conn, err := sql.Open("postgres", dbUrl)
	// if err != nil {
	// 	log.Fatal("Can't connect to database")
	// }

	router := chi.NewRouter()

	// Basic CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	// v1Router.Post("/users", apiCfg.handlerCreateUser)
	// v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUser))

	// v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	// v1Router.Get("/feeds", apiCfg.getFeeds)

	// v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handleGetPostsForUser))

	// v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	// v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.getUserFeedFollows))
	// v1Router.Delete("/feed_follows/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.deleteUserFeedFollow))

	router.Mount("/v1", v1Router)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	// Start serving an http server
	fmt.Printf("Server starting on port %v\n", port)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
