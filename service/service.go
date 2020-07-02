package service

import (
	"github.com/ArturMartini/go-demo-login-jwt/canonical"
	"github.com/ArturMartini/go-demo-login-jwt/jwt"
	"github.com/ArturMartini/go-demo-login-jwt/repository"
	"log"
	"sync"
)

type Service interface {
	Login(login canonical.Login) (*canonical.Jwt, error)
	Demo() error
}
type service struct {
	repo repository.Repository
}

var (
	once     = sync.Once{}
	instance Service
)

func New() Service {
	once.Do(func() {
		if instance == nil {
			instance = &service{
				repo: repository.New(),
			}
		}
	})
	return instance
}

func (r service) Login(login canonical.Login) (*canonical.Jwt, error) {
	user, err := r.repo.GetUser(login)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return jwt.Encode(user.Id)
}


func (r service) Demo() error {
	return nil
}