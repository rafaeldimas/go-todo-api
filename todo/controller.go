package todo

import (
	"net/http"
)

type todoController struct{}

func NewController() todoController {
	return todoController{}
}

func (*todoController) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create"))
}

func (*todoController) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List"))
}

func (*todoController) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update"))
}

func (*todoController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
