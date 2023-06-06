package helper

import (
	"errors"

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
