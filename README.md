# Authentication server and client using JWT

This project simulation the server jwt to generate token with private key and consumer decode and validate the token

##### Note: We have a web server to just simulate call endpoints

### Generate token
```go
//jwt.go
//When up the application execute initKeys for read in memory keys using to encode and decode JWT
//When call POST /login execute Encode function for adding metadata and generate JWT 
//With JWT on hands, call POST /client with http header "Authorization: Bearer $Access_Token
//When receive the request thw middleware check if token is valid using decode function
//Note: your keys should stay out of application files
package jwt

import (
	"crypto/rsa"
	"fmt"
	"github.com/ArturMartini/go-demo-login-jwt/canonical"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"time"
)

var (
	rsaPvtKey *rsa.PrivateKey
	rsaPbcKey *rsa.PublicKey
)

func init() {
	initKeys()
}

func initKeys() {
	pemPvt, err := ioutil.ReadFile("/secrets/keys-private.pem")
	if err != nil {
		log.Panicf("private key not found. err: %v", err)
	}

	pemPbc, err := ioutil.ReadFile("/secrets/keys-public.pem")
	if err != nil {
		log.Panicf("public key not found, err: %v", err)
	}

	keyPvt, err := jwt.ParseRSAPrivateKeyFromPEM(pemPvt)
	if err != nil {
		log.Panicf("cannot convert pem to private key, err: %v", err)
	}

	keyPbc, err := jwt.ParseRSAPublicKeyFromPEM(pemPbc)
	if err != nil {
		log.Panicf("cannot convert pem to public key, err: %v", err)
	}

	rsaPvtKey = keyPvt
	rsaPbcKey = keyPbc

	log.Printf("initialize jwt keys successfully")
}

func Encode(userId string) (*canonical.Jwt, error) {
	now := time.Now().UTC()
	expires := now.Add(time.Hour * 1)

	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.StandardClaims{
		Audience:  "http://yourdomain.com",
		ExpiresAt: expires.Unix(),
		IssuedAt:  now.Unix(),
		Subject:   userId,
	}

	tokenStr, err := token.SignedString(rsaPvtKey)
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

func Decode(hash string) (*jwt.Token, error) {
	token, err := jwt.Parse(hash, func(token *jwt.Token)(interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return rsaPbcKey, nil
	})

	return token, err
}
```

## For run 
*   Goto root project
*   go run main.go
*   curl -XPOST -d '{"username":"any", "password":"any"}' http://localhost:8080/login
*   curl -XPOST -H "Authorization: Bearer $access_token" http://localhost:8080/client

