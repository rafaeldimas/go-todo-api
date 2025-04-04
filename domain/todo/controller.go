package todo

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type TodoController interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	service TodoService
}

func NewController(service TodoService) TodoController {
	return &todoController{
		service: service,
	}
}

func (controller *todoController) Create(w http.ResponseWriter, r *http.Request) {
	todo := Todo{}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	newTodo, err := controller.service.Create(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func (controller *todoController) List(w http.ResponseWriter, r *http.Request) {
	todos, err := controller.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func (controller *todoController) Update(w http.ResponseWriter, r *http.Request) {
	todo := Todo{}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	updatedTodo, err := controller.service.Update(todo.Id, &todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTodo)
}

func (controller *todoController) Delete(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(queryId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := controller.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
