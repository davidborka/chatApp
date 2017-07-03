package middleware

import (
	"time"

	"github.com/davidborka/chatApp/api/auth"
	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateAuthToken(clientLogin string) string {
	expireToken := time.Now().Add(time.Minute * 25).Unix()
	claims := Claims{
		clientLogin,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:9000",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString(auth.SigningKey)
	return signedToken

}
