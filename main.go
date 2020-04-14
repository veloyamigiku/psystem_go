package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/veloyamigiku/psystem/internal/auth"
	"github.com/veloyamigiku/psystem/internal/data_type"
	"github.com/veloyamigiku/psystem/internal/db"
)

func main() {

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "443"
	}

	http.HandleFunc("/psystem/issue_jwt_for_signup", handleIssueJwtForSignup)
	http.HandleFunc("/psystem/signup", handleSignup)

	http.HandleFunc("/psystem/issue_jwt_for_point_operator_signup", handleIssueJWTForPointOperatorSignup)
	http.HandleFunc("/psystem/signup_for_point_operator", handleSignupForPointOperator)

	http.HandleFunc("/psystem/issue_jwt_for_user_login", handleIssueJwtForUserLogin)
	http.HandleFunc("/psystem/login", handleLogin)

	http.HandleFunc("/psystem/point/current", handleCurrentPoint)
	http.HandleFunc("/psystem/point/log", handlePointLog)

	http.HandleFunc("/psystem/issue_jwt_for_point_operator_login", handleIssueJwtForPointOperatorLogin)
	http.HandleFunc("/psystem/login_for_point_operator", handleLoginForPointOperator)

	http.HandleFunc("/psystem/point/add_history", handleAddPointHistory)

	// (create key)openssl genrsa -out https.key 2048
	// (create crt)openssl req -new -x509 -sha256 -key https.key -out https.crt -days 3650
	err := http.ListenAndServeTLS(
		":"+serverPort,
		"https.crt",
		"https.key",
		nil)
	if err != nil {
		panic(err)
	}

}

func response(w http.ResponseWriter, jsonStruct interface{}) {

	responseBytes, _ := json.Marshal(jsonStruct)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBytes)

}

// リクエストエラーを出力する。
func internalError(w http.ResponseWriter, errorJSON string) {
	errOutput := ([]byte)(errorJSON)
	w.Header().Set("Content-Type", "application/json")
	w.Write(errOutput)
}

// リクエストハンドラ（ポイント利用履歴を返却する）。
func handlePointLog(w http.ResponseWriter, r *http.Request) {

	// リクエストメソッドを確認する。
	if r.Method != http.MethodGet {
		internalError(w, `{"Log": null}`)
		return
	}

	// ポイント利用履歴を取得する。

	// ポイント利用履歴を出力する。
	resultPointLog := ([]byte)(
		`{"Log": {"AllCount": 1234,"FirstIndex": 0,"Count": 20,"1": {"Date": "2020/1/20 19:52:10","Detail": "shop_A","Point": 10},"2": {"Date": "2020/1/19 09:13:54","Detail": "shop_B","Point": 5}}}`)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultPointLog)
}

// リクエストハンドラ（ポイント残高を返却する）。
func handleCurrentPoint(w http.ResponseWriter, r *http.Request) {

	// リクエストメソッドを確認する。
	if r.Method != http.MethodGet {
		internalError(w, `{"Point": null}`)
		return
	}

	// ポイント残高を取得する。

	// ポイント残高を出力する。
	resultCurrentPoint := ([]byte)(`{"Point": 1234}`)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultCurrentPoint)
}

// リクエストハンドラ（利用者ログイン用_トークン発行）
func handleIssueJwtForUserLogin(w http.ResponseWriter, r *http.Request) {

	result := data_type.ResultIssueUserLoginJWT{
		Result: false,
	}

	if r.Method != http.MethodGet {
		response(w, result)
		return
	}

	token := auth.IssueJwt(5)
	result.Result = true
	result.Token = token

	response(w, result)

}

// リクエストハンドラ（ログイン処理）。
func handleLogin(w http.ResponseWriter, r *http.Request) {

	resultLogin := data_type.ResultLogin{
		Result: false,
		Token:  "",
	}

	token := r.Header.Get("Authorization")
	if !auth.ValidateJwt(token) {
		response(w, resultLogin)
	}

	// リクエストメソッドを確認する。
	if r.Method != http.MethodPost {
		response(w, resultLogin)
		return
	}

	// ログイン処理

	// リクエストボディを取得する。
	body, err := getRequestJSON(r)
	if err != nil {
		fmt.Println(err)
		response(w, resultLogin)
		return
	}

	// リクエストボディをオブジェクトに変換する。
	var login data_type.PostLogin
	json.Unmarshal(body, &login)
	paramName := login.Name
	paramPassword := login.Password

	// ユーザ名でデータベースを検索する。
	user, err := db.SearchUser(paramName)
	if err != nil {
		response(w, resultLogin)
		return
	}
	passwordHash := user.Password

	// パスワードをハッシュ化する。
	paramPasswordHash := auth.FromStringToMD5(paramPassword)

	// データベースのハッシュと上記のハッシュを比較する。
	if passwordHash != paramPasswordHash {
		response(w, resultLogin)
		return
	}

	// JWTトークンを発行する。
	newToken := auth.IssueJwt(30)

	// ログイン処理結果を出力する。
	resultLogin.Token = newToken
	resultLogin.Result = true
	response(w, resultLogin)
}

// リクエストハンドラ（ポイント操作元ログイン用_トークン発行）
func handleIssueJwtForPointOperatorLogin(w http.ResponseWriter, r *http.Request) {

	result := data_type.ResultIssuePointOperatorLoginJWT{
		Result: false,
	}

	if r.Method != http.MethodGet {
		response(w, result)
		return
	}

	token := auth.IssueJwt(5)
	result.Result = true
	result.Token = token

	response(w, result)

}

// リクエストハンドラ（ログイン処理_ポイント操作元）
func handleLoginForPointOperator(w http.ResponseWriter, r *http.Request) {

	result := data_type.ResultLoginForPointOperator{
		Result: false,
	}

	if r.Method != http.MethodPost {
		response(w, result)
		return
	}

	token := r.Header.Get("Authorization")
	if !auth.ValidateJwt(token) {
		response(w, result)
		return
	}

	body, err := getRequestJSON(r)
	if err != nil {
		response(w, result)
		return
	}

	var pointOperator data_type.PointOperator
	json.Unmarshal(body, &pointOperator)

	searchPointOperator, err := db.SearchPointOperator(pointOperator.Name)
	if err != nil {
		fmt.Println(err)
		response(w, result)
		return
	}

	hashedPassword := auth.FromStringToMD5(pointOperator.Password)
	if hashedPassword != searchPointOperator.Password {
		fmt.Print("password compare error.")
		response(w, result)
		return
	}

	newToken := auth.IssueJwt(30)
	result.Result = true
	result.Token = newToken

	response(w, result)

}

// リクエストハンドラ（利用者登録用_トークン発行）
func handleIssueJwtForSignup(w http.ResponseWriter, r *http.Request) {

	registerJwt := data_type.ResultIssueRegisterJWT{
		Result: false,
	}

	if r.Method != http.MethodGet {
		response(w, registerJwt)
		return
	}

	tokenString := auth.IssueJwt(5)
	registerJwt.Result = true
	registerJwt.Token = tokenString

	response(w, registerJwt)
}

// リクエストハンドラ（利用者登録）。
func handleSignup(w http.ResponseWriter, r *http.Request) {

	registerResult := data_type.ResultRegister{
		Result: false,
	}

	// リクエストメソッドを確認する。
	if r.Method != http.MethodPost {
		response(w, registerResult)
		return
	}

	fmt.Println("signup")

	// リクエストボディを取得する。
	body, err := getRequestJSON(r)
	if err != nil {
		fmt.Println(err)
		response(w, registerResult)
		return
	}

	// リクエストボディをオブジェクトに変換する。
	var user data_type.PostUser
	json.Unmarshal(body, &user)

	/*
		// リクエスト本体からJSONオブジェクトを取得する。
		postJSON, err := getPostJSON(r)
		if err != nil {
			fmt.Println(err)
			response(w, registerResult)
			return
		}
	*/

	// debug code
	paramName := user.Name
	paramPassword := user.Password
	paramUsername := user.UserName
	fmt.Printf("paramName:%s\n", paramName)
	fmt.Printf("paramPassword:%s\n", paramPassword)
	fmt.Printf("paramUsername:%s\n", paramUsername)

	// リクエストヘッダ（Authorization）を取得する。
	authorization := r.Header.Get("Authorization")
	fmt.Printf("Authorization:%s\n", authorization)
	// JWTトークンを検証する。
	if !auth.ValidateJwt(authorization) {
		fmt.Println(fmt.Errorf("JWT is invalid"))
		response(w, registerResult)
		return
	}

	// 登録処理
	err = db.Register(
		paramName,
		paramUsername,
		paramPassword)
	if err != nil {
		fmt.Println(err)
		response(w, registerResult)
		return
	}

	// 登録処理の結果を出力する。
	registerResult.Result = true
	response(w, registerResult)
}

// リクエストハンドラ（ポイント操作元登録用_トークン発行）
func handleIssueJWTForPointOperatorSignup(w http.ResponseWriter, r *http.Request) {

	result := data_type.ResultIssueRegisterJWTForPointOperator{
		Result: false,
	}

	if r.Method != http.MethodGet {
		response(w, result)
		return
	}

	token := auth.IssueJwt(5)
	result.Result = true
	result.Token = token

	response(w, result)

}

// リクエストハンドラ（ポイント操作元登録）。
func handleSignupForPointOperator(w http.ResponseWriter, r *http.Request) {

	result := data_type.ResultRegisterForPointOperator{
		Result: false,
	}

	if r.Method != http.MethodPost {
		response(w, result)
		return
	}

	token := r.Header.Get("Authorization")
	if !auth.ValidateJwt(token) {
		response(w, result)
	}

	fmt.Println("signup for Point Operator")

	body, err := getRequestJSON(r)
	if err != nil {
		response(w, result)
		return
	}

	var pointOperator data_type.PostPointOperator
	json.Unmarshal(body, &pointOperator)

	err = db.RegisterPointOperator(pointOperator)
	if err != nil {
		response(w, result)
		return
	}

	result.Result = true
	response(w, result)

}

func handleAddPointHistory(w http.ResponseWriter, r *http.Request) {

	resultAddPointHistory := data_type.ResultAddPointHistory{
		Result: false,
		Count:  0,
	}

	// リクエストメソッドを確認する。
	if r.Method != http.MethodPost {
		response(w, resultAddPointHistory)
		return
	}

	// リクエストボディ（JSON文字列）を取得する。
	body, err := getRequestJSON(r)
	if err != nil {
		response(w, resultAddPointHistory)
		return
	}

	// リクエストボディをオブジェクトに変換する。
	var pointAdds []data_type.PostPointHistory
	json.Unmarshal(body, &pointAdds)

	// ポイント操作情報を保存する。
	addCount, err := db.AddPointHistory(pointAdds)
	if err != nil {
		resultAddPointHistory.Result = false
	} else {
		resultAddPointHistory.Result = true
	}
	resultAddPointHistory.Count = addCount

	response(w, resultAddPointHistory)
}

func getRequestJSON(r *http.Request) (body []byte, err error) {

	// リクエストヘッダ（Content-Type）をチェックする。
	if r.Header.Get("Content-Type") != "application/json" {
		err = fmt.Errorf("Content-Type is not application/json.")
		return
	}

	// リクエストヘッダ（Content-Length）を取得する。
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		return
	}

	// リクエスト本体を取得する。
	body = make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		return
	}

	err = nil
	return
}
