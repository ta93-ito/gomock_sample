package main

import (
	"github.com/ta93-ito/gomock_sample/infra/route"
)

func main() {
	if err := route.Init(); err != nil {
		panic(err)
	}
}
