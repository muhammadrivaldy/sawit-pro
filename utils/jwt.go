package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJWT(userID int, signingKey string) (string, error) {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, &jwt.RegisteredClaims{
		ID:        fmt.Sprint(userID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
	})

	token, err := jwtToken.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return token, nil

}

func ParseJWT(token string, signingKey string) (int, error) {

	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil || !jwtToken.Valid {
		return 0, errors.New("jwt is not valid")
	}

	claims := jwtToken.Claims.(*jwt.RegisteredClaims)
	userID, err := strconv.ParseInt(claims.ID, 0, 32)

	return int(userID), err

}
