package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/veloyamigiku/psystem/internal/auth"
	"github.com/veloyamigiku/psystem/internal/db"
)

type RegisterJwt struct {
	Token  string `json:"token"`
	Result bool   `json:"result"`
}

type RegisterResult struct {
	Result bool `json:"result"`
}

func main() {

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	server := http.Server{
		Addr: ":" + serverPort,
	}
	http.HandleFunc("/psystem/issue_jwt_for_signup", handleIssueJwtForSignup)
	http.HandleFunc("/psystem/signup", handleSignup)
	http.HandleFunc("/psystem/login", handleLogin)
	http.HandleFunc("/psystem/point/current", handleCurrentPoint)
	http.HandleFunc("/psystem/point/log", handlePointLog)
	server.ListenAndServe()

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

// リクエストハンドラ（ログイン処理）。
func handleLogin(w http.ResponseWriter, r *http.Request) {

	// リクエストメソッドを確認する。
	if r.Method != http.MethodPost {
		internalError(w, `{"Result": false, "Token": null}`)
		return
	}

	// ログイン処理

	// ログイン処理結果を出力する。
	resultLogin := ([]byte)(`{"Result": true, "Token": "hogehoge"}`)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultLogin)
}

// リクエストハンドラ（利用者登録用_トークン発行）
func handleIssueJwtForSignup(w http.ResponseWriter, r *http.Request) {

	registerJwt := RegisterJwt{
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

	registerResult := RegisterResult{
		Result: false,
	}

	// リクエストメソッドを確認する。
	if r.Method != http.MethodPost {
		response(w, registerResult)
		return
	}

	fmt.Println("signup")

	// リクエスト本体からJSONオブジェクトを取得する。
	postJSON, err := getPostJSON(r)
	if err != nil {
		fmt.Println(err)
		response(w, registerResult)
		return
	}
	// debug code
	paramName := postJSON["name"].(string)
	paramPassword := postJSON["password"].(string)
	paramUsername := postJSON["username"].(string)
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

// リクエスト本体（JSON文字列）をJSONオブジェクトに変換して返却する。
func getPostJSON(r *http.Request) (
	postJSON map[string]interface{},
	err error) {

	// リクエストヘッダをチェックする。
	// Content-Typeが"application/json"であることを確認する。
	if r.Header.Get("Content-Type") != "application/json" {
		return nil, fmt.Errorf("Content-Type is not application/json.")
	}

	// リクエストヘッダをチェックする。
	// Content-Lengthが数値であること確認する。
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		return nil, err
	}

	// リクエスト本体を取得する。
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		return nil, err
	}

	// リクエスト本体をJSONオブジェクト(map[string]interface)に変換する。
	err = json.Unmarshal(body[:length], &postJSON)
	if err != nil {
		return nil, err
	}

	return postJSON, nil
}
