package auth

import (
	ue "basic-coding-kulina/modules/entity/user"

	"gorm.io/gorm"
)

type AuthRepo interface {
	GetUserByEmail(email string) (*ue.User, error)
	Login(email string) (*ue.AuthResponse, string, string, error)
	CreateUser(user *ue.RegisterRequest) error
	UserRecovery(userId string, codeVer string) error
	UpdateUserRecovery(userId string, codeVer string) error
	GetUserRecovery(userId string) (ue.UserRecovery, error)
	ChangePassword(user ue.RecoveryRequest) error
	DeleteUserRecovery(userId string) error
}

type authRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) AuthRepo {
	return &authRepo{
		db,
	}
}
