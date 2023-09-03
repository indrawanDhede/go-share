package helper

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(value string) (string, error) {
	claims := jwt.MapClaims{
		"sub": value,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Waktu kadaluwarsa token (contoh: 1 jam)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("your-secret-key") // Ganti dengan kunci rahasia yang kuat

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
