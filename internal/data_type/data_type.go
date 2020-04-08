package data_type

// ResultIssueRegisterJWT 利用者登録用のトークン発行結果の構造体。
type ResultIssueRegisterJWT struct {
	Token  string `json:"token"`
	Result bool   `json:"result"`
}

// ResultRegister 利用者登録結果の構造体。
type ResultRegister struct {
	Result bool `json:"result"`
}

// ResultLogin ログイン結果の構造体。
type ResultLogin struct {
	Result bool   `json:"result"`
	Token  string `json:"token"`
}

// ResultAddPointHistory ポイント操作情報登録結果の構造体。
type ResultAddPointHistory struct {
	// ポイント操作情報の登録結果。
	Result bool `json:"result"`
	// ポイント操作情報の登録件数。
	Count int `json:"count"`
}

// PostUser 送信利用者情報の構造体。
type PostUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	UserName string `json:"username"`
}

// PostLogin 送信ログイン情報の構造体。
type PostLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// PostPointHistory 送信ポイント操作情報の構造体。
type PostPointHistory struct {
	UserID int    `json:"user_id"`
	Date   int    `json:"date"`
	Detail string `json:"detail"`
	Point  int    `json:"point"`
}

// User usersテーブル用の構造体。
type User struct {
	ID       int    `gorm:"PRIMARY_KEY" gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"NOT NULL" gorm:"UNIQUE"`
	Username string `gorm:"NOT NULL"`
	Password string `gorm:"NOT NULL"`
}

// PointHistory point_historiesテーブル用の構造体。
type PointHistory struct {
	ID     int    `gorm:"PRIMARY_KEY" gorm:"AUTO_INCREMENT"`
	UserID int    `gorm:"NOT NULL"`
	Date   int    `gorm:"NOT NULL"`
	Detail string `gorm:"NOT NULL"`
	Point  int    `gorm:"NOT NULL"`
}
