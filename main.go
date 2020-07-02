package main

import (
	"github.com/ArturMartini/go-demo-login-jwt/middleware"
	"log"
)

func main() {
	log.Panic(middleware.Start())
}
