package seed

import (
	re "basic-coding-kulina/modules/entity/role"
)

func CreateRoles() []*re.Role {
	roles := []*re.Role{
		{Role: "Admin"},
		{Role: "User"},
	}

	return roles
}
