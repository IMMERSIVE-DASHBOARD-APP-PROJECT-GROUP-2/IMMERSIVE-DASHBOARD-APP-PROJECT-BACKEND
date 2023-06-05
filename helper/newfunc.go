package helper

import (
	"github.com/DASHBOARDAPP/features/user"
)

// Konstruktor untuk tipe UserRole
func NewUserRole(role string) user.UserRole {
	return user.UserRole(role)
}

// Konstruktor untuk tipe UserTeam
func NewUserTeam(team string) user.UserTeam {
	return user.UserTeam(team)
}
