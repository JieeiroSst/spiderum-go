package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/Spide_IT/spide_it/config"
	"time"
)

type TokenUser struct {
	conf *config.Config
}

func NewTokenUser(conf *config.Config) *TokenUser{
	return &TokenUser{conf:conf}
}

func (t *TokenUser) GenerateToken(username string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"]= username
	atClaims["exp"] = time.Now().Add(time.Hour * 60 * 60 *60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(t.conf.Secret.JwtSecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (t *TokenUser) ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.conf.Secret.JwtSecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "Missing Authentication Token", err
	}
}
