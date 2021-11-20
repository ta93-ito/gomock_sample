package usecase

import "github.com/ta93-ito/gomock_sample/domain/repository"

type UserUsecase struct {
	repo repository.User
}

// repoにmock repositoryを渡す
func NewUserUsecase(repo repository.User) *UserUsecase {
	return &UserUsecase{repo: repo}
}
