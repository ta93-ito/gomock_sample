package route

import (
	"net/http"

	"github.com/ta93-ito/gomock_sample/injector"
)

func Init() error {
	mux := http.NewServeMux()
	mux.Handle("/user", injector.NewUserHandler())
	err := http.ListenAndServe(":1234", mux)
	if err != nil {
		return err
	}
	return nil
}
