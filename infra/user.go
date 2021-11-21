package infra

import (
	"github.com/ta93-ito/gomock_sample/domain/entity"
	"github.com/ta93-ito/gomock_sample/domain/repository"
)

type userRepository struct{}

func NewUserRepository() repository.User {
	return &userRepository{}
}

func (ur *userRepository) Update(user *entity.User) error {
	return nil
}
