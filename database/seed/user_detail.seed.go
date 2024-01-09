package seed

import (
	ue "basic-coding-kulina/modules/entity/user"
)

func CreateUserDetail() []*ue.UserDetail {
	userDetail := []*ue.UserDetail{
		{
			Name:         "Administrator",
			Point:        0,
			Phone:        "08917283129283",
			ProfilePhoto: "https://storage.googleapis.com/ecowave/img/users/profile/profile.png",
			UserId:       1,
		},
		{
			Name:         "User 1",
			Point:        0,
			Phone:        "08917283109283",
			ProfilePhoto: "https://storage.googleapis.com/ecowave/img/users/profile/profile.png",
			UserId:       2,
		},
		{
			Name:         "User 2",
			Point:        0,
			Phone:        "0851728392716",
			ProfilePhoto: "https://storage.googleapis.com/ecowave/img/users/profile/profile.png",
			UserId:       3,
		},
	}

	return userDetail
}
