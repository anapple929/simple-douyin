package core

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"tokenutils/service"
)

type TokenService struct {
}

var jwtSecret = []byte("1122233")

type Claims struct {
	Id int64 `json:"id"`
	jwt.StandardClaims
}

func (*TokenService) GetIdByToken(ctx context.Context, req *service.GetIdByTokenRequest, out *service.GetIdByTokenResponse) error {
	token := req.UserToken
	token = string(token)

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) { return jwtSecret, nil })
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			out.UserId = int32(claims.Id)
			return nil

		}
	}
	return err

}
