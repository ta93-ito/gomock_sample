package usecase

import (
	"testing"

	"github.com/ta93-ito/gomock_sample/domain/entity"
	"github.com/ta93-ito/gomock_sample/mock_repository"

	"github.com/golang/mock/gomock"
)

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUser(ctrl)
	userUsecase := NewUserUsecase(mockRepo)
	param := &entity.User{
		ID:   1,
		Name: "hoge",
	}

	mockRepo.EXPECT().Update(param).Return(nil)

	err := userUsecase.Update(param)
	if err != nil {
		t.Errorf(err.Error())
	}
}
