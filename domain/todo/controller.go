package todo

import (
	"encoding/json"
	"log"
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
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTodo, err := controller.service.Create(&todo)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func (controller *todoController) List(w http.ResponseWriter, r *http.Request) {
	todos, err := controller.service.List()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func (controller *todoController) Update(w http.ResponseWriter, r *http.Request) {
	todo := Todo{}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	pathId := r.PathValue("id")
	id, err := strconv.Atoi(pathId)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTodo, err := controller.service.Update(id, &todo)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTodo)
}

func (controller *todoController) Delete(w http.ResponseWriter, r *http.Request) {
	pathId := r.PathValue("id")
	id, err := strconv.Atoi(pathId)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := controller.service.Delete(id); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
