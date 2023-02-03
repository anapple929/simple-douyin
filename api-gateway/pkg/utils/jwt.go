package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("1122233")

type Claims struct {
	Id int64 `json:"id"`
	jwt.StandardClaims
}

// 签发用户token
func GenerateToken(id int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * time.Hour)
	claims := Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "1122233",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) { return jwtSecret, nil })
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {

			return claims, nil
		}
	}
	return nil, err
}
