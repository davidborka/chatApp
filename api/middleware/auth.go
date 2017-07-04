package middleware

import (
	"fmt"
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
	fmt.Println(auth.SigningKey)
	key, err := jwt.ParseRSAPrivateKeyFromPEM(auth.SigningKey)
	if err != nil {
		fmt.Println(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	fmt.Println(key)
	signedToken, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err)

		fmt.Println("THE TOKEN IS:" + signedToken)

	}
	return signedToken
}
