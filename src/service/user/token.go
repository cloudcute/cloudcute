package user

import (
	"cloudcute/src/models/user"
	"cloudcute/src/pkg/token"
)

func CreateToken(u user.User) (string, error) {
	var t = token.Token{
		UserID: string(u.ID),
		Name: u.UserName,
	}
	return token.CreateToken(t)
}

func RefreshToken(tokenString string) (string, error) {
	return token.RefreshToken(tokenString)
}
