package main

import (
	"fmt"
	"net/http"
	"os"
)

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

	// 登録処理

	// 登録処理の結果を出力する。
	output := ([]byte)(`{"Result": true}`)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
