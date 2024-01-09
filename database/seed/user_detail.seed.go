package seed

import (
	ue "basic-coding-kulina/modules/entity/user"
)

func CreateUserDetail() []*ue.UserDetail {
	userDetail := []*ue.UserDetail{
		{
			Name:         "Supplier",
			Point:        0,
			Phone:        "08917283129283",
			ProfilePhoto: "assets/img/account.png",
			UserId:       "6c8bcb83-a825-4df0-8c29-264402205b9b",
		},
		{
			Name:         "User 1",
			Point:        0,
			Phone:        "08917283109283",
			ProfilePhoto: "assets/img/account.png",
			UserId:       "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
		},
		{
			Name:         "User 2",
			Point:        0,
			Phone:        "0851728392716",
			ProfilePhoto: "assets/img/account.png",
			UserId:       "41fb3d71-33bc-4a6e-9620-2d56f3090981",
		},
	}

	return userDetail
}
