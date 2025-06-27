package main

import (
	"log"
	"net/http"

	"custom-api-server/db"
	"custom-api-server/handlers"
	"custom-api-server/models"

	"github.com/gorilla/mux"
)

func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	db.Connect()
	db.DB.AutoMigrate(&models.User{})

	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/update", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/delete", handlers.DeleteUser).Methods("DELETE")
	// r.HandleFunc("/users/delete-all", handlers.DeleteAllUsers).Methods("DELETE")

	handlerWithCORS := withCORS(r)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlerWithCORS))
}
