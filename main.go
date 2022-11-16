package main

import (
	"fmt"
	"github.com/MszSabab/health/httpserver"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	Serve()
}

func Serve() {
	handler := GetHandler()
	err := http.ListenAndServe(fmt.Sprintf(":%d", 8000), handler)
	if err != nil {
		fmt.Println(err)
	}
}

// GetHandler - api
func GetHandler() http.Handler {
	router := chi.NewRouter()
	router.Get("/health", httpserver.Health)
	return router
}
