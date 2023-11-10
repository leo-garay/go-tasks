package main

import (
	"fmt"
	"leo-garay/go-tasks/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// Index Routes
	router.HandleFunc("/", handlers.IndexRoute)

	// Tasks Routes 
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.GetOneTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	
	router.HandleFunc("/jobs", handlers.GetJobs).Methods("GET")
	router.HandleFunc("/jobs", handlers.Create).Methods("POST")
	router.HandleFunc("/jobs/error", handlers.CreateError).Methods("POST")
	fmt.Println("Server started on port ", 3000)
	log.Fatal(http.ListenAndServe(":3000", router))
}