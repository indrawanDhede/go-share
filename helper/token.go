package helper

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Waktu kadaluwarsa token (contoh: 1 jam)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("key-token-rahasia") // Ganti dengan kunci rahasia yang kuat

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	secretKey := []byte("key-token-rahasia")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Metode tanda tangan tidak valid")
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}