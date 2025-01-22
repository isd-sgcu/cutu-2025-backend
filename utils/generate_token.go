package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

// generateTokens creates an access and refresh token
func GenerateTokens(userID string, jwtSecret string) (string, error) {
	// Access Token
	accessTokenClaims := jwt.MapClaims{
		"userId": userID,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	access, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return access, nil
}
