package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ta93-ito/gomock_sample/domain/entity"

	"github.com/ta93-ito/gomock_sample/usecase"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

type UpdateRequest struct {
	User entity.User
}

func (uh *UserHandler) Update(w http.ResponseWriter, r http.Request) {
	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := uh.usecase.Update(&req.User); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
