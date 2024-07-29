package container

import (
	"hotel.com/services"
)

type Services struct {
	services.ExampleServicer
	Other services.ExampleServicer
}

func Bind() *Services {
	return &Services{
		ExampleServicer: services.NewTestExampleService(),
	}
}
