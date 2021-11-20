package usecase

import (
	"testing"

	"github.com/ta93-ito/gomock_sample/domain/entity"
	"github.com/ta93-ito/gomock_sample/mock_repository"

	"github.com/golang/mock/gomock"
)

func TestUpdate(t *testing.T) {
	t.Run("id and name", func(t *testing.T) {
		testUpdate(t, &entity.User{ID: 1, Name: "hoge"})
	})
	t.Run("id only", func(t *testing.T) {
		testUpdate(t, &entity.User{ID: 2})
	})
	t.Run("name only", func(t *testing.T) {
		testUpdate(t, &entity.User{Name: "hoge"})
	})
}

func testUpdate(t *testing.T, in *entity.User) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUser(ctrl)
	userUsecase := NewUserUsecase(mockRepo)

	mockRepo.EXPECT().Update(in).Return(nil)

	err := userUsecase.Update(in)
	t.Error()
	if err != nil {
		t.Errorf(err.Error())
	}

}
