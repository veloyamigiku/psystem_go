package data_type

type ResultLogin struct {
	Result bool   `json:"result"`
	Token  string `json:"token"`
}

type ResultIssueRegisterJWT struct {
	Token  string `json:"token"`
	Result bool   `json:"result"`
}

type ResultRegister struct {
	Result bool `json:"result"`
}

type ResultPointAdd struct {
	// ポイント加算情報の登録結果。
	Result bool `json:"result"`
	// ポイント加算情報の登録件数。
	Count int `json:"count"`
}
