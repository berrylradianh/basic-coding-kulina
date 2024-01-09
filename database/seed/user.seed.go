package seed

import (
	"basic-coding-kulina/helper/password"
	ue "basic-coding-kulina/modules/entity/user"
	"time"
)

func CreateUser() []*ue.User {
	hashPasswordUser1, _ := password.HashPassword("user1")
	hashPasswordUser2, _ := password.HashPassword("user2")
	hashPasswordSupplier, _ := password.HashPassword("supplier123")
	users := []*ue.User{
		{
			ID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Email:     "supplier@gmail.com",
			Username:  "supplier",
			Password:  string(hashPasswordSupplier),
			RoleId:    "419a8a2d-0abe-4413-ac49-39d33cf9838d",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
			Email:     "user1@gmail.com",
			Username:  "user1",
			Password:  string(hashPasswordUser1),
			RoleId:    "b29112fc-c7b5-4386-a31e-f2c040de7fcb",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "41fb3d71-33bc-4a6e-9620-2d56f3090981",
			Email:     "user2@gmail.com",
			Username:  "user2",
			Password:  string(hashPasswordUser2),
			RoleId:    "b29112fc-c7b5-4386-a31e-f2c040de7fcb",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return users
}
