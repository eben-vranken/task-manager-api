package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eben-vranken/task-manager-api/internal/database"
	"github.com/eben-vranken/task-manager-api/internal/handlers"
	"github.com/eben-vranken/task-manager-api/internal/repository"
)

const PORT string = "8080"

var databaseURL = os.Getenv("DATABASE_URL")

func main() {
	http.HandleFunc("GET /health", healthCheck)

	db, err := database.New(databaseURL)

	taskRepository := &repository.TaskRepository{DB: db}
	taskHandler := &handlers.TaskHandler{TR: taskRepository}

	if err != nil {
		log.Panicf("Error when opening database %v\n", err)
	}

	http.HandleFunc("POST /task", taskHandler.Create)

	fmt.Println("Listening to port", PORT+"...")
	log.Fatal(http.ListenAndServe("127.0.0.1:"+PORT, nil))
}

func healthCheck(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Up and running!"))
}
