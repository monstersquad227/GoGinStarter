package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

/*
	Encrypt 使用 AES GCM 算法加密数据
*/

func Encrypt(data, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	seal := gcm.Seal(nonce, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(seal), nil
}

/*
	Decrypt 使用 AES GCM 算法解密数据
*/

func Decrypt(data, key string) (string, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(dataByte) < nonceSize {
		return "", errors.New("dataByte too short")
	}
	nonce, ciphertext := dataByte[:nonceSize], dataByte[nonceSize:]
	open, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(open), nil
}
