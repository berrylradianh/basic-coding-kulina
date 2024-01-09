package seed

import (
	re "basic-coding-kulina/modules/entity/role"
	"time"
)

func CreateRoles() []*re.Role {
	roles := []*re.Role{
		{
			ID:        "419a8a2d-0abe-4413-ac49-39d33cf9838d",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Role:      "Supplier",
		},
		{
			ID:        "b29112fc-c7b5-4386-a31e-f2c040de7fcb",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Role:      "User",
		},
	}

	return roles
}
