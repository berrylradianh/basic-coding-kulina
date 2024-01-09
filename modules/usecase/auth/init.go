package auth

import (
	ue "basic-coding-kulina/modules/entity/user"
	ar "basic-coding-kulina/modules/repository/auth"
)

type AuthUsecase interface {
	Login(request *ue.LoginRequest) (interface{}, uint, error)
	Register(user *ue.RegisterRequest) error
	ForgotPassword(request ue.ForgotPassRequest) (string, error)
	VerifOtp(request ue.VerifOtp) error
	ChangePassword(request ue.RecoveryRequest) error
}

type authUsecase struct {
	authRepo ar.AuthRepo
}

func New(adminRepo ar.AuthRepo) *authUsecase {
	return &authUsecase{
		adminRepo,
	}
}
