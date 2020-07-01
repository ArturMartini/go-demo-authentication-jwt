package repository

import "authentication/src/authentication/canonical"

type Repository interface {
	GetUser(login canonical.Login) (canonical.User, error)
}

type repo struct {}

func (repo) GetUser(login canonical.Login) (canonical.User, error) {
	return canonical.User{
		Id: "1",
	}, nil
}

func New() Repository {
	return new(repo)
}
