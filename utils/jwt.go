package utils

import (
	"crypto/rsa"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func getPrivateKey() (*rsa.PrivateKey, error) {

	privateKey, err := os.ReadFile("../private/jwt.key")
	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPrivateKeyFromPEM(privateKey)

}

func getPublicKey() (*rsa.PublicKey, error) {

	publicKey, err := os.ReadFile("../private/jwt-public.key")
	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPublicKeyFromPEM(publicKey)

}

func CreateJWT(userId int) (string, error) {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id":  userId,
		"exp": jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		"jti": uuid.New(),
	})

	privateKey, err := getPrivateKey()
	if err != nil {
		return "", err
	}

	token, err := jwtToken.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return token, nil

}

func ParseJWT(token string) (int, error) {

	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "RS256" {
			return nil, errors.New("unexpected signing method")
		}

		return getPublicKey()
	})
	if err != nil || !jwtToken.Valid {
		return 0, errors.New("jwt is not valid")
	}

	claims := jwtToken.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)

	return int(userId), err

}
