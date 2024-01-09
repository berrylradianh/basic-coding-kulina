package auth

import (
	ue "basic-coding-kulina/modules/entity/user"

	"gorm.io/gorm"
)

type AuthRepo interface {
	GetUserByEmail(email string) (*ue.User, error)
	Login(email string) (*ue.AuthResponse, string, string, error)
	CreateUser(user *ue.RegisterRequest) error
}

type authRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) AuthRepo {
	return &authRepo{
		db,
	}
}
