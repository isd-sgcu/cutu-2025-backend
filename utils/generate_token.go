package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// generateTokens creates an access and refresh token
func GenerateTokens(userID string, role string, jwtSecret string) (string, string, error) {
	// Access Token
	accessTokenClaims := jwt.MapClaims{
		"userId": userID,
		"role":   role,
		"exp":    time.Now().Add(45 * time.Minute).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	access, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", "", err
	}

	// Refresh Token
	refreshTokenClaims := jwt.MapClaims{
		"userId": userID,
		"role":   role,
		"exp":    time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refresh, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}
