package repository

import "github.com/ta93-ito/gomock_sample/domain/entity"

type User interface {
	Update(user *entity.User) error
}
