package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/eben-vranken/task-manager-api/internal/models"
	"github.com/eben-vranken/task-manager-api/internal/repository"
	"github.com/jackc/pgx/v5/pgconn"
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

func (uh *UserHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	users, err := uh.ur.GetAll(req.Context())

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)

	if err != nil {
		log.Print(err)
		log.Println("500 - Internal Server error")
	}
}

func (uh *UserHandler) GetSpecificById(w http.ResponseWriter, req *http.Request) {
	user, err := uh.ur.GetSpecificById(req.Context(), req.PathValue("id"))

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal server error"))
		return
	}

	if *user == (models.User{}) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - User not found"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		log.Print(err)
		log.Print("500 - Internal server error")
	}
}

func (uh *UserHandler) Delete(w http.ResponseWriter, req *http.Request) {
	result, err := uh.ur.Delete(req.Context(), req.PathValue("id"))

	var constraintError *pgconn.PgError
	if errors.As(err, &constraintError) {
		log.Print(err)
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("409 - Cannot delete user since they have tasks assigned to them"))
		return
	}

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

func (uh *UserHandler) Update(w http.ResponseWriter, req *http.Request) {
	var user models.User

	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("200 - Bad request"))
		return
	}

	newUser, err := uh.ur.Update(req.Context(), user, req.PathValue("id"))

	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - User not found"))
		return
	}

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(newUser)

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
