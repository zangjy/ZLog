package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

const (
	ivLength = 16
)

// EncryptString 使用AES算法加密数据，返回Base64编码的加密结果
func EncryptString(data string, key string) (string, error) {
	if len(data) == 0 || len(key) == 0 {
		return "", errors.New("data or key is empty")
	}

	iv, err := generateRandomIV()
	if err != nil {
		return "", err
	}

	adjustedKey := adjustKeySize(key)

	block, err := aes.NewCipher([]byte(adjustedKey))
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, len(data))

	cipher.NewCFBEncrypter(block, iv).XORKeyStream(cipherText, []byte(data))

	combined := append(iv, cipherText...)
	return base64.StdEncoding.EncodeToString(combined), nil
}

// EncryptBytes 使用AES算法加密数据，返回加密后的字节数组
func EncryptBytes(data []byte, key string) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New("data or key is empty")
	}

	iv, err := generateRandomIV()
	if err != nil {
		return nil, err
	}

	adjustedKey := adjustKeySize(key)

	block, err := aes.NewCipher([]byte(adjustedKey))
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, len(data))

	cipher.NewCFBEncrypter(block, iv).XORKeyStream(cipherText, data)

	return append(iv, cipherText...), nil
}

// DecryptString 使用AES算法解密Base64编码的加密数据，返回解密后的字符串
func DecryptString(data string, key string) (string, error) {
	if len(data) == 0 || len(key) == 0 {
		return "", errors.New("data or key is empty")
	}

	combined, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	iv := combined[:ivLength]
	cipherText := combined[ivLength:]

	adjustedKey := adjustKeySize(key)

	block, err := aes.NewCipher([]byte(adjustedKey))
	if err != nil {
		return "", err
	}

	plainText := make([]byte, len(cipherText))

	cipher.NewCFBDecrypter(block, iv).XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}

// DecryptBytes 使用AES算法解密字节数组，返回解密后的字节数组
func DecryptBytes(data []byte, key string) ([]byte, error) {
	if len(data) == 0 || len(key) == 0 {
		return nil, errors.New("data or key is empty")
	}

	iv := data[:ivLength]
	cipherText := data[ivLength:]

	adjustedKey := adjustKeySize(key)

	block, err := aes.NewCipher([]byte(adjustedKey))
	if err != nil {
		return nil, err
	}

	plainText := make([]byte, len(cipherText))

	cipher.NewCFBDecrypter(block, iv).XORKeyStream(plainText, cipherText)

	return plainText, nil
}

// adjustKeySize 处理密钥长度，确保它是16、24或32字节长
func adjustKeySize(key string) string {
	keyLength := len(key)

	if keyLength == 16 || keyLength == 24 || keyLength == 32 {
		return key
	} else if keyLength < 16 {
		return key + string(make([]byte, 16-keyLength))
	} else if keyLength < 24 {
		return key + string(make([]byte, 24-keyLength))
	} else {
		return key[:32]
	}
}

// generateRandomIV 生成随机的初始化向量 (IV)
func generateRandomIV() ([]byte, error) {
	iv := make([]byte, ivLength)
	_, err := io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}
	return iv, nil
}
