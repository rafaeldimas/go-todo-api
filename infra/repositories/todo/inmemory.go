package todo

import (
	"errors"

	domain "github.com/rafaeldimas/go-todo-api/domain/todo"
)

type inMemoryTodoRepository struct {
	todos []*domain.Todo
}

func NewInMemoryRepository() domain.TodoRepository {
	return &inMemoryTodoRepository{
		todos: make([]*domain.Todo, 0),
	}
}

func (repository *inMemoryTodoRepository) Create(todo *domain.Todo) (*domain.Todo, error) {
	todo.Id = len(repository.todos) + 1
	repository.todos = append(repository.todos, todo)

	return todo, nil
}

func (repository *inMemoryTodoRepository) List() ([]*domain.Todo, error) {
	return repository.todos, nil
}

func (repository *inMemoryTodoRepository) Update(id int, todo *domain.Todo) (*domain.Todo, error) {
	for i, t := range repository.todos {
		if t.Id == id {
			repository.todos[i] = todo

			return todo, nil
		}
	}

	return &domain.Todo{}, errors.New("todo not found")
}

func (repository *inMemoryTodoRepository) Delete(id int) error {
	for i, t := range repository.todos {
		if t.Id == id {
			repository.todos = append(repository.todos[:i], repository.todos[i+1:]...)

			return nil
		}
	}

	return errors.New("todo not found")
}
