package auth

import (
	"errors"

	re "basic-coding-kulina/modules/entity/role"
	ue "basic-coding-kulina/modules/entity/user"

	"github.com/google/uuid"
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

	UserID := uuid.New().String()

	userTable := ue.User{
		ID:       UserID,
		RoleId:   role.ID,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		UserDetail: ue.UserDetail{
			UserId: UserID,
			Name:   user.Name,
			Phone:  user.Phone,
		},
	}

	if err := ar.db.Create(&userTable).Error; err != nil {
		return err
	}
	return nil
}
