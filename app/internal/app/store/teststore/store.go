package teststore

import (
	"github.com/DmitryOdintsov/rest-api/app/internal/app/model"
	"github.com/DmitryOdintsov/rest-api/app/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
