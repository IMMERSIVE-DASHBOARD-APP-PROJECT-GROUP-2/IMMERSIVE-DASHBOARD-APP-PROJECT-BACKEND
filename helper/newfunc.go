package helper

import (
	"errors"
	"unicode"

	"github.com/DASHBOARDAPP/app/config"
	"github.com/DASHBOARDAPP/features/user"
	"github.com/golang-jwt/jwt"
)

// // Konstruktor untuk tipe UserRole
// func NewUserRole(role string) user.UserRole {
// 	return user.UserRole(role)
// }

// // Konstruktor untuk tipe UserTeam
//
//	func NewUserTeam(team string) user.UserTeam {
//		return user.UserTeam(team)
//	}
func validateTokenAndGetRole(token string) (user.UserRole, error) {
	// Validasi token JWT
	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET_JWT), nil
	})
	if err != nil {
		return "", errors.New("invalid token")
	}

	// Memeriksa apakah token telah diverifikasi
	if !parsedToken.Valid {
		return "", errors.New("invalid token")
	}

	// Mendapatkan role pengguna dari token
	role := claims["role"].(string)

	return user.UserRole(role), nil
}

// Fungsi helper untuk memeriksa apakah string hanya terdiri dari angka
func containsOnlyNumbers(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// Fungsi helper untuk memeriksa apakah string mengandung simbol
func containsSymbols(s string) bool {
	for _, r := range s {
		if unicode.IsSymbol(r) || unicode.IsPunct(r) {
			return true
		}
	}
	return false
}
