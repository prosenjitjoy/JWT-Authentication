package main

import (
	"context"
	"log"
	"net/http"
	"server/database"
	"server/internal/user"
	"server/middlewares"
	"server/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	ctx := context.Background()
	conn, err := database.NewDatabase(ctx)
	if err != nil {
		log.Fatal("Could not initialize database connection")
	}
	defer conn.Close(ctx)

	userAdapter := user.NewAdapter(conn.GetDB())
	userService := user.NewService(userAdapter)
	userHandler := user.NewHandler(userService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middlewares.Cors)
	routes.Use(router, userHandler)

	err = http.ListenAndServe(":5000", router)
	if err != nil {
		log.Fatal("Could not initialize http server")
	}
}
