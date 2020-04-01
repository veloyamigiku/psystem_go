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

// ResultPointAdd ポイント加算結果の構造体。
type ResultPointAdd struct {
	// ポイント加算情報の登録結果。
	Result bool `json:"result"`
	// ポイント加算情報の登録件数。
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

// PostPointAdd 送信ポイント加算情報の構造体。
type PostPointAdd struct {
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
