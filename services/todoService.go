package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-sample-api/dto"
	"golang-sample-api/models"
	"golang-sample-api/repositories"
)

//go:generate mockgen -destination=../mocks/service/mockTodoService.go -package=services golang-sample-api/services TodoService
type DefaultTodoService struct {
	Repo repositories.TodoRepository
}

type TodoService interface {
	TodoInsert(todo models.Todo) (*dto.TodoDto, error)
	TodoGetAll() ([]models.Todo, error)
	TodoDelete(id primitive.ObjectID) (bool, error)
}

func (t DefaultTodoService) TodoInsert(todo models.Todo) (*dto.TodoDto, error) {
	var res dto.TodoDto
	if len(todo.Title) <= 2 {
		res.Status = false
		return &res, nil
	}
	result, err := t.Repo.Insert(todo)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}
	res = dto.TodoDto{Status: result}
	return &res, nil
}

func (t DefaultTodoService) TodoGetAll() ([]models.Todo, error) {
	result, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t DefaultTodoService) TodoDelete(id primitive.ObjectID) (bool, error) {
	result, err := t.Repo.Delete(id)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func NewTodoService(Repo repositories.TodoRepository) DefaultTodoService {
	return DefaultTodoService{Repo: Repo}
}
