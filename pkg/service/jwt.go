package service

import (
	"coinkeeper/configs"
	"coinkeeper/logger"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// CustomClaims определяет кастомные поля токена
type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Role     string `json:"role"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken генерирует JWT токен с кастомными полями
func GenerateToken(userID uint, username string, role string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Duration(configs.AppSettings.AuthParams.JwtTtlMinutes) * time.Minute), // токен истекает через 1 час
			Issuer:    configs.AppSettings.AppParams.ServerName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

// ParseToken парсит JWT токен и возвращает кастомные поля
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи токена
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		logger.Error.Println("[service.ParseToken] cannot parse token. Error is:", err.Error())
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	logger.Error.Println("[service.ParseToken] invalid token")
	return nil, fmt.Errorf("invalid token")
}

/*

 */
