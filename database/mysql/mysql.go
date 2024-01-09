package mysql

import (
	"fmt"

	"basic-coding-kulina/config"
	"basic-coding-kulina/database/seed"
	pe "basic-coding-kulina/modules/entity/product"
	re "basic-coding-kulina/modules/entity/role"
	et "basic-coding-kulina/modules/entity/transaction"
	ue "basic-coding-kulina/modules/entity/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	InitDB()
	InitialMigration()
	seed.DBSeed(DB)
}
func InitDB() {
	var err error

	configurations := config.GetConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configurations.DB_USERNAME,
		configurations.DB_PASSWORD,
		configurations.DB_HOST,
		configurations.DB_PORT,
		configurations.DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
}

func InitialMigration() {
	DB.AutoMigrate(
		re.Role{},
		ue.User{},
		ue.UserDetail{},
		ue.UserRecovery{},
		ue.UserAddress{},
		pe.ProductCategory{},
		pe.Product{},
		pe.ProductImage{},
		et.Transaction{},
		et.TransactionDetail{},
	)
	DB.Migrator().HasConstraint(&ue.User{}, "UserDetail")
	DB.Migrator().HasConstraint(&re.Role{}, "Users")
	DB.Migrator().HasConstraint(&pe.Product{}, "TransactionDetail")
}
