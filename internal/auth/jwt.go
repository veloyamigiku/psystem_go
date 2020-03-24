package auth

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func IssueJwt(expMinute int) (tokenString string) {

	// 秘密鍵の読み込み。
	// ssh-keygen -t rsa
	sign := os.Getenv("PSYSTEM_RSA")
	if sign == "" {
		panic(fmt.Errorf("env[PSYSTEM_RSA] not found"))
	}
	signBytes := []byte(sign)
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

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

func ValidateJwt(tokenString string) bool {

	// 公開鍵の読み込み。
	// ssh-keygen -f (pub_key_name) -e -m pkcs8 > (pub_key_name).pkcs8
	verify := os.Getenv("PSYSTEM_RSA_PUB_PKCS8")
	if verify == "" {
		panic(fmt.Errorf("env[PSYSTEM_RSA_PUB_PKCS8] not found"))
	}
	verifyByte := []byte(verify)
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
