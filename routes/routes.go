package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/olagookundavid/rssagg/handlers"
)

func CreateRouter(apiCfg *handlers.ApiConfig) chi.Router {
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
	v1Router.Get("/healthz", handlers.HandlerReadiness)
	v1Router.Get("/err", handlers.HandlerErr)

	v1Router.Post("/users", apiCfg.HandlerCreateUser)
	v1Router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandleGetUser))

	v1Router.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.GetFeeds)

	v1Router.Get("/posts", apiCfg.MiddlewareAuth(apiCfg.HandleGetPostsForUser))

	v1Router.Post("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.GetUserFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowId}", apiCfg.MiddlewareAuth(apiCfg.DeleteUserFeedFollow))

	router.Mount("/v1", v1Router)
	return router
}
