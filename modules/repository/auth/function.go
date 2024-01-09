package auth

import (
	"errors"

	re "basic-coding-kulina/modules/entity/role"
	ue "basic-coding-kulina/modules/entity/user"
)

func (ar *authRepo) GetUserByEmail(email string) (*ue.User, error) {
	user := &ue.User{}
	err := ar.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, errors.New("Record Not Found")
	}

	return user, nil
}
func (ar *authRepo) Login(email string) (*ue.AuthResponse, string, string, error) {
	user := &ue.User{}
	err := ar.db.Preload("UserAddresses").Preload("UserDetail").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, "", "", errors.New("Record Not Found")
	}

	var address ue.UserAddress
	for _, val := range user.UserAddresses {
		if val.IsPrimary == true {
			address = val
		}
	}

	response := &ue.AuthResponse{
		ID:            user.ID,
		Email:         user.Email,
		Username:      user.Username,
		Name:          user.UserDetail.Name,
		Point:         user.UserDetail.Point,
		Phone:         user.UserDetail.Phone,
		ProfilePhoto:  user.UserDetail.ProfilePhoto,
		UserAddresses: address,
	}

	return response, user.Password, user.RoleId, nil
}

func (ar *authRepo) CreateUser(user *ue.RegisterRequest) error {
	existingUser := ue.User{}
	if err := ar.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return errors.New("Email already exists")
	}

	var role re.Role
	if err := ar.db.Where("role = ?", "User").First(&role).Error; err != nil {
		return err
	}

	userTable := ue.User{
		RoleId:   role.Role,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		UserDetail: ue.UserDetail{
			Name:  user.Name,
			Phone: user.Phone,
		},
	}

	if err := ar.db.Create(&userTable).Error; err != nil {
		return err
	}
	return nil
}
func (ar *authRepo) GetUserRecovery(userId string) (ue.UserRecovery, error) {
	var recovery ue.UserRecovery
	err := ar.db.Where("user_id = ?", userId).First(&recovery).Error
	if err != nil {
		return recovery, errors.New("Record Not Found")
	}

	return recovery, nil
}

func (ar *authRepo) UserRecovery(userId string, codeVer string) error {

	userRecover := ue.UserRecovery{
		UserId: userId,
		Code:   codeVer,
	}
	if err := ar.db.Create(&userRecover).Error; err != nil {
		return err
	}

	return nil
}
func (ar *authRepo) UpdateUserRecovery(userId string, codeVer string) error {

	userRecover := ue.UserRecovery{
		UserId: userId,
		Code:   codeVer,
	}
	result := ar.db.Model(&ue.UserRecovery{}).Where("user_id = ?", userId).Updates(&userRecover)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return errors.New("nothing has changed")
	}

	return nil
}
func (ar *authRepo) ChangePassword(user ue.RecoveryRequest) error {

	result := ar.db.Model(&ue.User{}).Where("email = ?", user.Email).Update("password", user.Password)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return errors.New("nothing has changed")
	}

	return nil
}
func (ar *authRepo) DeleteUserRecovery(userId string) error {

	var userRecovery ue.UserRecovery
	result := ar.db.Where("user_id = ?", userId).Delete(&userRecovery)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return errors.New("nothing has changed")
	}

	return nil
}
