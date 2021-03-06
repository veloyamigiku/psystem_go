package db

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/veloyamigiku/psystem/internal/auth"
	"github.com/veloyamigiku/psystem/internal/data_type"
)

// Register 利用者を登録する。
func Register(name string, username string, password string) error {

	// パスワードをハッシュ化する。
	passwordHashStr := auth.FromStringToMD5(password)

	dbRes := db.Create(&data_type.User{
		Name:     name,
		Username: username,
		Password: passwordHashStr,
	})

	return dbRes.Error

}

// SearchUser 利用者を検索する。
func SearchUser(name string) (user data_type.User, err error) {

	dbRes := db.First(&user, "name = ?", name)
	err = dbRes.Error

	return

}

// SearchPointOperator ポイント操作元を検索する。
func SearchPointOperator(name string) (pointOperator data_type.PointOperator, err error) {

	dbRes := db.First(&pointOperator, "name = ?", name)
	err = dbRes.Error

	return
}

// RegisterPointOperator ポイント操作元を登録する。
func RegisterPointOperator(pointOperator data_type.PostPointOperator) (err error) {

	hashed_password := auth.FromStringToMD5(pointOperator.Password)

	dbRes := db.Create(&data_type.PointOperator{
		Name:     pointOperator.Name,
		Password: hashed_password,
	})
	err = dbRes.Error

	return
}

// AddPointHistory ポイント操作情報を登録する。
func AddPointHistory(pointAdds []data_type.PostPointHistory) (addCount int, err error) {

	addCount = 0
	for _, pointAdd := range pointAdds {
		dbRes := db.Create(&data_type.PointHistory{
			UserID: pointAdd.UserID,
			Date:   pointAdd.Date,
			Detail: pointAdd.Detail,
			Point:  pointAdd.Point,
		})
		if dbRes.Error != nil {
			return
		}
		addCount++
	}

	return
}
