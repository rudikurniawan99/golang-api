package helper

import (
	"os"
	"strconv"

	jwt "github.com/golang-jwt/jwt/v4"
)

type jwtClaim struct {
	id string
	jwt.RegisteredClaims
}

func GenerateToken(id int) (string, error) {
	strId := strconv.Itoa(id)
	claim := jwtClaim{
		strId,
		jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(os.Getenv("PUBLIC_KEY")))

	if err != nil {
		return "", err
	} else {
		return signedToken, err
	}

}
