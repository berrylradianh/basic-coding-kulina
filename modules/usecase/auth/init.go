package auth

import (
	ue "basic-coding-kulina/modules/entity/user"
	ar "basic-coding-kulina/modules/repository/auth"
)

type AuthUsecase interface {
	Login(request *ue.LoginRequest) (interface{}, string, error)
	Register(user *ue.RegisterRequest) error
}

type authUsecase struct {
	authRepo ar.AuthRepo
}

func New(adminRepo ar.AuthRepo) *authUsecase {
	return &authUsecase{
		adminRepo,
	}
}
