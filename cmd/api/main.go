package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/eben-vranken/task-manager-api/internal/database"
	"github.com/eben-vranken/task-manager-api/internal/handlers"
	"github.com/eben-vranken/task-manager-api/internal/repository"
)

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.statusCode = code
	sr.ResponseWriter.WriteHeader(code)
}

const PORT string = "8080"

var databaseURL = os.Getenv("DATABASE_URL")

func main() {
	http.HandleFunc("GET /health", loggingMiddleware(healthCheck))

	db, err := database.New(databaseURL)

	if err != nil {
		log.Panicf("Error when opening database %v\n", err)
	}

	taskRepository := repository.NewTaskRepository(db)
	taskHandler := handlers.NewTaskHandler(taskRepository)

	http.HandleFunc("POST /task", loggingMiddleware(taskHandler.Create))
	http.HandleFunc("GET /task", loggingMiddleware(taskHandler.GetAll))
	http.HandleFunc("GET /task/{id}", loggingMiddleware(taskHandler.GetSpecificById))
	http.HandleFunc("DELETE /task/{id}", loggingMiddleware(taskHandler.Delete))
	http.HandleFunc("PUT /task/{id}", taskHandler.Update)
	fmt.Println("Listening to port", PORT+"...")
	log.Fatal(http.ListenAndServe("127.0.0.1:"+PORT, nil))
}

func healthCheck(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Up and running!"))
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.URL.Path, "Initializing logging middleware")
		start := time.Now()

		recorder := &statusRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(recorder, req)
		duration := time.Since(start)
		log.Printf("[%s] %s %s %d", req.Method, req.RequestURI, duration, recorder.statusCode)
	})
}
