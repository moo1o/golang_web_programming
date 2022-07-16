package internal

import (
	"errors"
	"log"
)

type Application struct {
	repository Repository
}

var (
	ErrAlreadyExist = errors.New("Membership already exist")
)

func NewApplication(repository Repository) *Application {
	// repository.data
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	if _, ok := app.repository.data[request.UserName]; ok {
		log.Println(ErrAlreadyExist)
		return CreateResponse{}, ErrAlreadyExist
	}

	return CreateResponse{request.UserName, request.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	return UpdateResponse{}, nil
}

func (app *Application) Delete(id string) error {
	return nil
}
