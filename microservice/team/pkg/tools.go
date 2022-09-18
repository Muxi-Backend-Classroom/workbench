package pkg

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"workbench-team/config"
)

type Claims struct {
	UserId uint32
	jwt.StandardClaims
}

func GenerateUrl(id uint32) (string, error) {
	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			Issuer:    config.Cfg.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return tokenClaims.SignedString("Muxi")
}
