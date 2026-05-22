package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eben-vranken/task-manager-api/internal/models"
	"github.com/eben-vranken/task-manager-api/internal/repository"
)

type TaskHandler struct {
	tr *repository.TaskRepository
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

	createdTask, err := th.tr.Create(req.Context(), task)

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

func (th *TaskHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	tasks, err := th.tr.GetAll(req.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Interval server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tasks)

	if err != nil {
		log.Print(err)
		log.Println("500 - Interval server error")
	}
}

func (th *TaskHandler) GetSpecificById(w http.ResponseWriter, req *http.Request) {
	task, err := th.tr.GetSpecificById(req.Context(), req.PathValue("id"))

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Interval server error"))
		return
	}

	if *task == (models.Task{}) {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Task not found"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(task)

	if err != nil {
		log.Print(err)
		log.Println("500 - Internal server error")
	}
}

func (th *TaskHandler) Delete(w http.ResponseWriter, req *http.Request) {
	result, err := th.tr.Delete(req.Context(), req.PathValue("id"))

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	err = json.NewEncoder(w).Encode(result)

	if err != nil {
		log.Print(err)
		log.Println("500 - Internal server error")
	}
}

func (th *TaskHandler) Update(w http.ResponseWriter, req *http.Request) {
	var newTask models.Task

	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&newTask)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("200 - Bad Request"))
		return
	}

	task, err := th.tr.Update(req.Context(), newTask, req.PathValue("id"))

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(task)
}

func NewTaskHandler(tr *repository.TaskRepository) *TaskHandler {
	t := new(TaskHandler)
	t.tr = tr
	return t
}
