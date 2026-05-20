package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eben-vranken/task-manager-api/internal/models"
	"github.com/eben-vranken/task-manager-api/internal/repository"
)

type TaskHandler struct {
	TR *repository.TaskRepository
}

func (th *TaskHandler) Create(w http.ResponseWriter, req *http.Request) {
	var task models.Task

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&task)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad request"))
		return
	}

	createdTask, err := th.TR.Create(req.Context(), task)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdTask)

	if err != nil {
		log.Print(err)
		log.Println("500 - Interval server error")
	}
}
