package uaes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"
)

func en_aesgcm(hex_key string, plaintext []byte, hex_aad string) (string, error) {
	//秘钥长度按需:AES-128(16bytes)或者AES-256(32bytes)
	key, _ := hex.DecodeString(hex_key)
	aad, _ := hex.DecodeString(hex_aad)

	block, err := aes.NewCipher(key) //生成加解密用的block
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	// 对IV有随机性要求，但没有保密性要求，所以常见的做法是将IV包含在加密文本当中
	iv := make([]byte, aesgcm.NonceSize())            // NonceSize=12
	rand.Read(iv)                                     //获取随机值
	ciphertext := aesgcm.Seal(iv, iv, plaintext, aad) //加密,密文为:iv+密文+tag
	//return base64.RawStdEncoding.EncodeToString(ciphertext), nil // 生成的BS64,无尾部的pad"="
	return base64.StdEncoding.EncodeToString(ciphertext), nil // 生成的BS64
}
func de_aesgcm(hex_key string, cipherbs64 string, hex_aad string) ([]byte, error) {
	//秘钥长度按需:AES-128(16bytes)或者AES-256(32bytes)
	key, _ := hex.DecodeString(hex_key)
	aad, _ := hex.DecodeString(hex_aad)

	cipherbs64 = strings.TrimRight(cipherbs64, "=")
	ciphertext, err := base64.RawStdEncoding.DecodeString(cipherbs64) // 要先去掉尾部的pad"=",否则解bs64失败
	if err != nil {
		return []byte(""), err
	}
	block, err := aes.NewCipher(key) //生成加解密用的block
	if err != nil {
		return []byte(""), err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte(""), err
	}
	if len(ciphertext) <= aesgcm.NonceSize() { // 长度应该>iv
		return []byte(""), errors.New("string: too short") //解密失败
	}

	iv := ciphertext[:aesgcm.NonceSize()]        //分离出IV
	ciphertext = ciphertext[aesgcm.NonceSize():] // 密文+tag
	plaintext, err := aesgcm.Open(nil, iv, ciphertext, aad)
	return plaintext, err //err!=nil时,plaintext=byte[]("")
}
