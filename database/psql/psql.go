package psql

import (
	"fmt"
	"os"

	"basic-coding-kulina/database/seed"
	pe "basic-coding-kulina/modules/entity/product"
	re "basic-coding-kulina/modules/entity/role"
	et "basic-coding-kulina/modules/entity/transaction"
	ue "basic-coding-kulina/modules/entity/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	InitDB()
	InitialMigration()
	seed.DBSeed(DB)
}
func InitDB() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		username,
		password,
		name,
		host,
		port,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.Migrator().DropTable(
		&re.Role{},
		&ue.User{},
		&ue.UserDetail{},
		&ue.UserRecovery{},
		&ue.UserAddress{},
		&pe.ProductCategory{},
		&pe.Product{},
		&pe.ProductImage{},
		&et.Transaction{},
		&et.TransactionDetail{},
	)
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
