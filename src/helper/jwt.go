package helper

import (
	"os"
	"strconv"

	jwt "github.com/golang-jwt/jwt/v4"
)

type jwtClaim struct {
	ID string `json:"id"`
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

func ValidateToken(signedToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("PUBLIC_KEY")), nil
		},
	)

	claims, ok := token.Claims.(*jwtClaim)
	if ok && token.Valid {
		return claims.ID, nil
	} else {
		return claims.ID, err
	}

}
