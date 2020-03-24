package db

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/veloyamigiku/psystem/internal/auth"
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

	// パスワードをハッシュ化する。
	passwordHashStr := auth.FromStringToMD5(password)

	dbRes := db.Create(&User{
		Name:     name,
		Username: username,
		Password: passwordHashStr,
	})

	return dbRes.Error

}

// SearchUser 利用者を検索する。
func SearchUser(name string) (user User, err error) {

	dbRes := db.First(&user, "name = ?", name)
	err = dbRes.Error

	return

}
