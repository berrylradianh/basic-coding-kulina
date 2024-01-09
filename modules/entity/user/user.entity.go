package user

import (
	et "basic-coding-kulina/modules/entity/transaction"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"type:text;primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	RoleId        string
	Email         string           `json:"Email" form:"Email" validate:"required,email"`
	Username      string           `json:"Username" form:"Username" validate:"required"`
	Password      string           `json:"Password" form:"Password" validate:"required,min=8"`
	UserDetail    UserDetail       `gorm:"foreignKey:UserId"`
	UserAddresses []UserAddress    `gorm:"foreignKey:UserId"`
	Transactions  []et.Transaction `gorm:"foreignKey:UserId"`
}

type RegisterRequest struct {
	Name     string `json:"Name" form:"Name" validate:"required"`
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Username string `json:"Username" form:"Username" validate:"required"`
	Phone    string `json:"Phone" form:"Phone" validate:"required,min=10,max=15,numeric"`
	Password string `json:"Password" form:"Password" validate:"required,min=8"`
}
type LoginRequest struct {
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Password string `json:"Password" form:"Password" validate:"required"`
}

type AuthResponse struct {
	ID            string `json:"Id" form:"Id"`
	Email         string `json:"Email" form:"Email" validate:"required,email"`
	Username      string `json:"Username" form:"Username" validate:"required"`
	Name          string `json:"Name" form:"Name"`
	Point         uint   `json:"Point" form:"Point"`
	Phone         string `json:"Phone" form:"Phone" validate:"required,min=10,max=13"`
	ProfilePhoto  string `json:"ProfilePhoto" form:"ProfilePhoto"`
	UserAddresses UserAddress
	Token         string `json:"Token" form:"Token"`
}

type UserLogin struct {
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Password string `json:"Password" form:"Password" validate:"required"`
}

type UserResponseLogin struct {
	Email    string
	Username string
	Token    string
}

type UserRequest struct {
	Email    string `json:"Email" form:"Email" validate:"email"`
	Username string `json:"Username" form:"Username"`
}

type UserResponse struct {
	Id           string `json:"Id"`
	RoleId       string `json:"RoleId"`
	Name         string `json:"Name"`
	Username     string `json:"Username"`
	Email        string `json:"Email"`
	Phone        string `json:"Phone"`
	Point        uint   `json:"Point"`
	ProfilePhoto string `json:"ProfilePhoto"`
	Addresses    []UserAddressResponse
}
