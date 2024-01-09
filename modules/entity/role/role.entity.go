package role

import (
	ut "basic-coding-kulina/modules/entity/user"

	"gorm.io/gorm"
)

type Role struct {
	*gorm.Model

	Role  string `json:"Role" form:"Role"`
	Users []ut.User
}
