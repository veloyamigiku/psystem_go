package auth

import (
	"fmt"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func IssueJwt() (tokenString string) {

	// 秘密鍵の読み込み。
	signBytes, err := ioutil.ReadFile("/root/go/src/github.com/veloyamigiku/psystem/demo.rsa")
	if err != nil {
		panic(err)
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	// JWTトークンを発行する。
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "test"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err = token.SignedString(signKey)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func ValidateJwt(tokenString string) bool {

	// 公開鍵の読み込み。
	verifyByte, err := ioutil.ReadFile("/root/go/src/github.com/veloyamigiku/psystem/demo.rsa.pub.pkcs8")
	if err != nil {
		panic(err)
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyByte)
	if err != nil {
		panic(err)
	}

	// JWTトークンを検証する。
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, err := token.Method.(*jwt.SigningMethodRSA)
		if !err {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return verifyKey, nil
	})
	if err == nil && token.Valid {
		return true
	} else {
		return false
	}
}
