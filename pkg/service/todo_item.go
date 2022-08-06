package service

import (
	"todolist"
	"todolist/pkg/repo"
)

type TodoItemService struct {
	repo     repo.TodoItem
	listRepo repo.TodoList
}

func NewTodoItemService(repo repo.TodoItem, listRepo repo.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (s *TodoItemService) Create(userId, listId int, item *todolist.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	id, err := s.repo.Create(listId, item)

	return id, nil
}

func (s *TodoItemService) GetAll(userId, listId int) (*[]todolist.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (todolist.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, id int, input *todolist.UpdateItemInput) error {
	return s.repo.Update(userId, id, input)
}
