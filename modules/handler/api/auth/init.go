package auth

import ac "basic-coding-kulina/modules/usecase/auth"

type AuthHandler struct {
	authUsecase ac.AuthUsecase
}

func New(authUsecase ac.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase,
	}
}
