package auth

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// readPrivateKey 秘密鍵を読み込む。
func readPrivateKey() (signKey *rsa.PrivateKey, err error) {

	// ssh-keygen -t rsa
	sign := os.Getenv("PSYSTEM_RSA")
	if sign == "" {
		err = fmt.Errorf("env[PSYSTEM_RSA] not found")
		return
	}
	signBytes := []byte(sign)
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)

	return
}

// IssueJwtAfterLogin ログイン後のトークンを発行する。
func IssueJwtAfterLogin(expMinute int, name string) (tokenString string, err error) {

	// 秘密鍵を読み込む。
	singKey, err := readPrivateKey()
	if err != nil {
		return
	}

	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Duration(expMinute) * time.Minute).Unix()
	tokenString, err = token.SignedString(singKey)

	return
}

// IssueJwt トークンを発行する。
func IssueJwt(expMinute int) (tokenString string) {

	// 秘密鍵の読み込み。
	signKey, err := readPrivateKey()

	// JWTトークンを発行する。
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "test"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Duration(expMinute) * time.Minute).Unix()

	tokenString, err = token.SignedString(signKey)
	if err != nil {
		fmt.Println(err)
	}

	return
}

// ValidateJwt トークンを検証する。
func ValidateJwt(tokenString string) (token *jwt.Token, err error) {

	// 公開鍵の読み込み。
	// ssh-keygen -f (pub_key_name) -e -m pkcs8 > (pub_key_name).pkcs8
	verify := os.Getenv("PSYSTEM_RSA_PUB_PKCS8")
	if verify == "" {
		err = fmt.Errorf("env[PSYSTEM_RSA_PUB_PKCS8] not found")
		return
	}
	verifyByte := []byte(verify)
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyByte)
	if err != nil {
		return
	}

	// JWTトークンを検証する。
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, err := token.Method.(*jwt.SigningMethodRSA)
		if !err {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return verifyKey, nil
	})
	return
}
