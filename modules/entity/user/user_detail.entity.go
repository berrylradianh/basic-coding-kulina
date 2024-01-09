package user

import "gorm.io/gorm"

type UserDetail struct {
	*gorm.Model  `json:"-"`
	Name         string `json:"Name" form:"Name" validate:"required"`
	Point        uint   `json:"Point" form:"Point"`
	Phone        string `json:"Phone" form:"Phone" validate:"required,min=10,max=13,numeric"`
	ProfilePhoto string `json:"ProfilePhoto" form:"ProfilePhoto"`
	UserId       string `json:"UserId" form:"UserId"`
}

type UserDetailRequest struct {
	Name         string `json:"Name" form:"Name"`
	Phone        string `json:"Phone" form:"Phone" validate:"min=10,max=13"`
	ProfilePhoto string `json:"ProfilePhoto" form:"ProfilePhoto"`
}
