package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ta93-ito/gomock_sample/domain/entity"
	"github.com/ta93-ito/gomock_sample/mock_repository"
	"github.com/ta93-ito/gomock_sample/usecase"
)

func TestUpdate(t *testing.T) {
	t.Run("id and name", func(t *testing.T) {
		testUpdate(t, &UpdateRequest{entity.User{ID: 1, Name: "hoge"}})
	})
}

func testUpdate(t *testing.T, in *UpdateRequest) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUser(ctrl)
	mockRepo.EXPECT().Update(&in.User).Return(nil)

	userUsecase := usecase.NewUserUsecase(mockRepo)
	userHandler := NewUserHandler(userUsecase)

	reqBody, err := json.Marshal(in)
	if err != nil {
		t.Errorf(err.Error())
	}

	readBody := bytes.NewReader(reqBody)
	r := httptest.NewRequest("PUT", "/user", readBody)
	w := httptest.NewRecorder()

	userHandler.ServeHTTP(w, r)
	res := w.Result()
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Invalid Status Code")
	}
}
