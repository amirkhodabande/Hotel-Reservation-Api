package services

import "fmt"

type ExampleServicer interface {
	JustTest(value string) error
}

type TestExampleService struct {
}

func NewTestExampleService() *TestExampleService {
	return &TestExampleService{}
}

func (s *TestExampleService) JustTest(value string) error {
	fmt.Println("from TestExampleService implementation...")
	return nil
}
