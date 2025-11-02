package apptoken

import (
	"maps"

	"github.com/golang-jwt/jwt"
)

func tokenParse(tokenString string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
	return token, err
}

func Encript(data map[string]any, secretKey string) (string, error) {

	claim := jwt.MapClaims{}
	maps.Copy(claim, data)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Decript(tokenString string, secretKey string) (map[string]any, error) {

	token, err := tokenParse(tokenString, secretKey)
	if err != nil {
		return map[string]any{}, err
	}

	claims := token.Claims.(jwt.MapClaims)

	result := map[string]any{}
	maps.Copy(result, claims)
	return result, nil
}

func IsTokenValid(tokenString string, secretKey string) bool {
	token, err := tokenParse(tokenString, secretKey)
	if err != nil {
		return false
	}
	return token.Valid
}
