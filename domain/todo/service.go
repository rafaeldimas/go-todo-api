package todo

type TodoService interface {
	Create(todo *Todo) (*Todo, error)
	List() ([]*Todo, error)
	Update(id int, todo *Todo) (*Todo, error)
	Delete(id int) error
}

type todoService struct {
	repository TodoRepository
}

func NewService(repository TodoRepository) TodoService {
	return &todoService{
		repository: repository,
	}
}

func (service *todoService) Create(todo *Todo) (*Todo, error) {
	return service.repository.Create(todo)
}

func (service *todoService) List() ([]*Todo, error) {
	return service.repository.List()
}

func (service *todoService) Update(id int, todo *Todo) (*Todo, error) {
	return service.repository.Update(id, todo)
}

func (service *todoService) Delete(id int) error {
	return service.repository.Delete(id)
}
