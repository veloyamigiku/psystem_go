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
