package service

import (
	"authentication/src/authentication/canonical"
	"authentication/src/authentication/repository"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

type Service interface {
	Login(login canonical.Login) (*canonical.Jwt, error)
}
type service struct {
	repo repository.Repository
}

var (
	once = sync.Once{}
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
		return  nil, err
	}
	return jwtGenerate(user.Id)
}

func jwtGenerate(userId string) (*canonical.Jwt, error) {
	now := time.Now().UTC()
	expires := now.Add(time.Hour * 1)

	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.StandardClaims{
		Audience:  "http://yourdomain.com",
		ExpiresAt: expires.Unix(),
		IssuedAt:  now.Unix(),
		Subject:   userId,
	}


	key, err := ioutil.ReadFile("keys-private.pem")
	if err != nil {
		log.Println(err.Error())
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		log.Println(err.Error())
	}

	tokenStr, err := token.SignedString(signKey)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	return &canonical.Jwt{
		AccessToken: tokenStr,
		ExpiresIn:   expires.Unix(),
		TokenType:   "Bearer",
	}, nil
}

