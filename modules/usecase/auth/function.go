package auth

import (
	"errors"

	pw "basic-coding-kulina/helper/password"
	vld "basic-coding-kulina/helper/validator"
	"basic-coding-kulina/middleware/jwt"
	ue "basic-coding-kulina/modules/entity/user"
)

func (ac *authUsecase) Register(request *ue.RegisterRequest) error {
	if err := vld.Validation(request); err != nil {
		return err
	}

	hashedPassword, err := pw.HashPassword(request.Password)
	if err != nil {
		return err
	}

	request.Password = string(hashedPassword)

	err = ac.authRepo.CreateUser(request)
	if err != nil {
		return err
	}

	return nil
}

func (ac *authUsecase) Login(request *ue.LoginRequest) (interface{}, string, error) {
	if err := vld.Validation(request); err != nil {
		return nil, "", err
	}

	response, password, role, err := ac.authRepo.Login(request.Email)

	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return nil, "", errors.New("Email atau password salah")
	}

	err = pw.VerifyPassword(password, request.Password)
	if err != nil {

		//lint:ignore ST1005 Reason for ignoring this linter
		return nil, "", errors.New("Email atau password salah")
	}

	token, err := jwt.CreateToken(response.ID, response.Email)
	if err != nil {
		return nil, "", err
	}
	response.Token = token

	return response, role, nil
}
