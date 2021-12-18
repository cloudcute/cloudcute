package token

import (
	"cloudcute/src/pkg/config"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var (
	Expired     error  = errors.New("token is expired")
	NotValidYet error  = errors.New("token not active yet")
	Malformed   error  = errors.New("that's not even a token")
	Invalid     error  = errors.New("couldn't handle this token")
	Notfound    error  = errors.New("token notfound")
)

type Token struct {
	jwt.StandardClaims
	UserID  string `json:"userId"`
	Name    string `json:"name"`
}

func parseToken(signKey string, tokenString string) (*Token, error) {
	if tokenString == "" {
		return nil, Notfound
	}
	var key = []byte(signKey)
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			var err error
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				err =  Malformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				err =  Expired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				err =  NotValidYet
			} else {
				err =  Invalid
			}
			return nil, err
		}
	}else{
		if claims, ok := token.Claims.(*Token); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, Invalid
}

func createToken(signKey string, token Token) (string, error) {
	var key = []byte(signKey)
	var t = jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	return t.SignedString(key)
}

func refreshToken(signKey string, tokenString string) (string, error) {
	var key = []byte(signKey)
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Token); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return createToken(signKey, *claims)
	}
	return "", Invalid
}

func getSignKey() string {
	return config.SystemConfig.SignKey
}

func ParseToken(tokenString string) (*Token, error) {
	return parseToken(getSignKey(), tokenString)
}

func CreateToken(token Token) (string, error) {
	return createToken(getSignKey(), token)
}

func RefreshToken(tokenString string) (string, error) {
	return refreshToken(getSignKey(), tokenString)
}
