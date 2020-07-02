package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testTokenString string

func TestInit(t *testing.T) {
	initKeys()
	assert.NotNil(t, rsaPbcKey)
	assert.NotNil(t, rsaPvtKey)
}

func TestJwtEncode(t *testing.T) {
	jwt, err := Encode("123")
	assert.Nil(t, err)
	assert.NotEmpty(t, jwt.AccessToken)
	assert.NotEmpty(t, jwt.ExpiresIn)
	assert.NotEmpty(t, jwt.TokenType)
	testTokenString = jwt.AccessToken
}

func TestJwtDecode(t *testing.T) {
	token, err := Decode(testTokenString)
	claims := token.Claims.(jwt.MapClaims)
	iat := int64(claims["iat"].(float64))
	exp := int64(claims["exp"].(float64))
	now := time.Now().Unix()
	assert.Nil(t, err)
	assert.Equal(t, "123", claims["sub"])
	assert.Equal(t, "http://yourdomain.com", claims["aud"])
	assert.Equal(t, true, now >= iat)
	assert.Equal(t, true, now < exp)
}
