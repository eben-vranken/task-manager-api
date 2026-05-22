package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eben-vranken/task-manager-api/internal/models"
	"github.com/eben-vranken/task-manager-api/internal/repository"
)

type UserHandler struct {
	ur *repository.UserRepository
}

func (uh *UserHandler) Create(w http.ResponseWriter, req *http.Request) {
	var user models.User
	
	decoder := json.NewDecoder(req.Body)
	
	err := decoder.Decode(&user)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("200 - Bad request"))
		return
	}

	createdUser, err := uh.ur.Create(req.Context(), user) 

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdUser)

	if err != nil {
		log.Print(err)
		log.Println("500 - Internal server error")
	}
}

func NewUserHandler(ur *repository.UserRepository) *UserHandler {
	t := new(UserHandler)
	t.ur = ur
	return t
}