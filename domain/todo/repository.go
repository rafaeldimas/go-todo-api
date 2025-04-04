package todo

type TodoRepository interface {
	Create(todo *Todo) (*Todo, error)
	List() ([]*Todo, error)
	Update(id int, todo *Todo) (*Todo, error)
	Delete(id int) error
}
