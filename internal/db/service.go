package db

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User usersテーブルの構造体。
type User struct {
	ID       int    `gorm:"PRIMARY_KEY" gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"NOT NULL" gorm:"UNIQUE"`
	Username string `gorm:"NOT NULL"`
	Password string `gorm:"NOT NULL"`
}

// Register 利用者を登録する。
func Register(name string, username string, password string) error {

	dbRes := db.Create(&User{
		Name:     name,
		Username: username,
		Password: password,
	})

	return dbRes.Error

}
