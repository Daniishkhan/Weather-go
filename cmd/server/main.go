package main

import (
	"context"
	"fmt"
	"goweather/pkg/database"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.ConnectToDb(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to db %v", err)
	}
	defer db.Close(context.Background())

	fmt.Println("Connect to db successfully")

	fmt.Println("Starting server...")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", helloWorld)

	fmt.Println("Server listening on 3000")
	http.ListenAndServe(":3000", r)

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to Goweather"))
}
