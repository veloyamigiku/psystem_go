package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User usersテーブルの構造体。
type User struct {
	ID       int    `gorm:"PRIMARY_KEY";"AUTO_INCREMENT"`
	Name     string `gorm:"NOT NULL"`
	Username string `gorm:"NOT NULL"`
	Password string `gorm:"NOT NULL"`
}

func main() {

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	server := http.Server{
		Addr: ":" + serverPort,
	}
	http.HandleFunc("/psystem/signup", handleSignup)
	http.HandleFunc("/psystem/login", handleLogin)
	http.HandleFunc("/psystem/point/current", handleCurrentPoint)
	http.HandleFunc("/psystem/point/log", handlePointLog)
	server.ListenAndServe()

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

// リクエストハンドラ（利用者登録）。
func handleSignup(w http.ResponseWriter, r *http.Request) {

	// リクエストメソッドを確認する。
	if r.Method != http.MethodPost {
		internalError(w, `{"Result": false}`)
		return
	}

	fmt.Println("signup")

	// リクエスト本体からJSONオブジェクトを取得する。
	postJSON, err := getPostJSON(r)
	if err != nil {
		fmt.Println(err)
		internalError(w, `{"Result": false}`)
		return
	}
	// debug code
	user := postJSON["user"].(string)
	password := postJSON["password"].(string)
	username := postJSON["username"].(string)
	fmt.Printf("user:%s\n", user)
	fmt.Printf("password:%s\n", password)
	fmt.Printf("username:%s\n", username)

	// 登録処理

	// DBに接続する。
	db, err := gorm.Open(
		"postgres",
		"host=D8C74545-postgres port=5432 user=psystem dbname=psystem password=psystem sslmode=disable")
	if err != nil {
		panic(err)
		defer db.Close()
	}
	defer db.Close()
	// 利用者情報を登録する。
	db.Create(&User{
		Name:     user,
		Username: username,
		Password: password,
	})

	// 登録処理の結果を出力する。
	output := ([]byte)(`{"Result": true}`)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
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
