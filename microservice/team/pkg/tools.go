package pkg

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"workbench-team/config"
)

type Claims struct {
	Infor
	jwt.StandardClaims
}

type Infor struct {
	Id      uint32
	GroupId uint32
}

func GenerateUrl(uid uint32, groupId uint32) (string, error) {
	claims := Claims{
		Infor{Id: uid, GroupId: groupId},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			Issuer:    config.Cfg.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return tokenClaims.SignedString("Muxi")
}

func ParseUrl(url string) (Infor, error) {
	claims, err := jwt.ParseWithClaims(url, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.Issuer), nil
	})
	if err != nil || claims == nil {
		return Infor{}, err
	}
	if claim, ok := claims.Claims.(*Claims); ok {
		return claim.Infor, nil
	}
	return Infor{}, nil
}
