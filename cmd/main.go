package main

import (
	"log"
	"net/http"

	"github.com/xddprog/internal/core/repositories"
	"github.com/xddprog/internal/core/services"
	"github.com/xddprog/internal/handlers"
	"github.com/xddprog/internal/infrastructure/database/adapters"
)


func main() {
	db, err := adapters.NewSQLiteConnection()
	if err != nil {
		log.Fatal(err)
	}

	server := http.NewServeMux()

	quoteRepository := repositories.NewQuoteRepository(db)
	quoteService := services.NewQuoteService(quoteRepository)
	quoteHandler := handlers.NewQuoteHandler(quoteService)

	quoteHandler.SetupRoutes(server)
	http.ListenAndServe("localhost:8080", server)
}
