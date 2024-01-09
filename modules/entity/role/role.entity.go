package role

import (
	ut "basic-coding-kulina/modules/entity/user"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        string `gorm:"type:text;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Role      string         `json:"Role" form:"Role"`
	Users     []ut.User
}
