package models

import (
    "github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"time"
    jwt "github.com/dgrijalva/jwt-go"
	"strings"
	"fmt"
	//"strings"
	//"github.com/go-xorm/xorm"
)

var FilterUser = func(ctx *context.Context) {
	token := ctx.Input.Header("Authorization")
 
	b, _ := CheckToken(token)
 
	//Verify that Token is legal
	if !b {
		http.Error(ctx.ResponseWriter, "Token verification not pass", http.StatusBadRequest)
		return
	}
	
}

func Fatal(err error) {
	if err != nil {
		beego.Error(err)
	}
}

func CreateToken(username string) (string) {
  claims := make(jwt.MapClaims)
			claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
			claims["iat"] = time.Now().Unix()
			claims["nameid"] = username
			claims["User"] = "true"
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		 
			tokenString, _ := token.SignedString([]byte(beego.AppConfig.String("jwtkey")))
		 
			return tokenString

}

func CheckToken(token string) (b bool, t *jwt.Token) {
	kv := strings.Split(token, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		beego.Error("AuthString invalid:", token)
		return false, nil
	}
	t, err := jwt.Parse(kv[1], func(*jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("jwtkey")), nil
	})
	fmt.Println(t)
	if err != nil {
		fmt.Println("Convert to jwt claims fail.", err)
		return false, nil
	}
	return true, t
}



func Encrypt(text string) string {
    key := []byte(beego.AppConfig.String("secretkey"))
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

func Decrypt(cryptoText string) string {
    key := []byte(beego.AppConfig.String("secretkey"))
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
