package todo

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
	domain "github.com/rafaeldimas/go-todo-api/domain/todo"
)

type mysqlTodoRepository struct {
	db *sql.DB
}

func NewMysqlRepository() domain.TodoRepository {
	url := os.Getenv("MYSQL_DATABASE_URL")
	db, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	return &mysqlTodoRepository{
		db: db,
	}
}

func (repository *mysqlTodoRepository) Create(todo *domain.Todo) (*domain.Todo, error) {
	result, err := repository.db.Exec("INSERT INTO todos (name, done) VALUES (?, ?)", todo.Name, todo.Done)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	todo.Id = int(id)
	return todo, nil

}

func (repository *mysqlTodoRepository) List() ([]*domain.Todo, error) {
	rows, err := repository.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}

	todos := []*domain.Todo{}
	for rows.Next() {
		var todo domain.Todo
		rows.Scan(&todo.Id, &todo.Name, &todo.Done)
		todos = append(todos, &todo)
	}

	return todos, nil
}

func (repository *mysqlTodoRepository) Update(id int, todo *domain.Todo) (*domain.Todo, error) {
	result, err := repository.db.Exec("UPDATE todos SET name = ?, done = ? WHERE id = ?", todo.Name, todo.Done, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("todo not found")
	}

	return todo, nil
}

func (repository *mysqlTodoRepository) Delete(id int) error {
	_, err := repository.db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
