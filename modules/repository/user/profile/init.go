package profile

import (
	ut "basic-coding-kulina/modules/entity/user"

	"gorm.io/gorm"
)

type ProfileRepo interface {
	GetAllUserProfile(user *[]ut.User) error
	GetUserProfile(user *ut.User, id string) error
	GetUserDetailProfile(userDetail *ut.UserDetail, id string) error
	UpdateUserProfile(userRequest *ut.UserRequest, id string) error
	UpdateUserDetailProfile(userDetailRequest *ut.UserDetailRequest, id string) error
	CreateAddressProfile(address *ut.UserAddress) error
	GetAllAddressProfileNoPagination(address *[]ut.UserAddress, idUser string) error
	GetAllAddressProfile(address *[]ut.UserAddress, idUser string, offset, pageSize int) (*[]ut.UserAddress, int64, error)
	GetAddressByIdProfile(address *ut.UserAddress, idUser string, idAddress int) error
	UpdateAddressPrimaryProfile(address *ut.UserAddress, idUser string) error
	UpdateAddressByIdProfile(addressRequest *ut.UserAddressRequest, idUser string, idAddress int) error
	UpdatePasswordProfile(newPassword string, id string) error
}

type profileRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ProfileRepo {
	return &profileRepo{
		db,
	}
}
