package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// Encrypt 函数将输入数据使用AES算法加密，并返回Base64编码的加密结果。
func Encrypt(data string, key string) (string, error) {
	// 处理密钥长度
	adjustedKey := adjustKeySize(key)

	// 创建AES密码块
	block, err := aes.NewCipher([]byte(adjustedKey))
	if err != nil {
		return "", err
	}

	// 创建初始化向量（IV）
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 创建CFB加密器
	encrypter := cipher.NewCFBEncrypter(block, iv)

	// 加密数据
	encrypted := make([]byte, len(data))
	encrypter.XORKeyStream(encrypted, []byte(data))

	// 将IV与加密数据合并并返回Base64编码结果
	combined := append(iv, encrypted...)
	return base64.StdEncoding.EncodeToString(combined), nil
}

// Decrypt 函数将Base64编码的加密数据使用AES算法解密，并返回原始数据。
func Decrypt(data string, key string) (string, error) {
	// 处理密钥长度
	adjustedKey := adjustKeySize(key)

	// 解码Base64数据
	combined, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	// 提取IV和加密数据
	iv := combined[:aes.BlockSize]
	encrypted := combined[aes.BlockSize:]

	// 创建AES密码块
	block, err := aes.NewCipher([]byte(adjustedKey))
	if err != nil {
		return "", err
	}

	// 创建CFB解密器
	decrypter := cipher.NewCFBDecrypter(block, iv)

	// 解密数据
	decrypted := make([]byte, len(encrypted))
	decrypter.XORKeyStream(decrypted, encrypted)

	// 返回解密后的原始数据
	return string(decrypted), nil
}

// adjustKeySize 函数用于处理密钥长度，确保它是16、24或32字节长。
func adjustKeySize(key string) string {
	keyLength := len(key)
	switch {
	case keyLength == 16, keyLength == 24, keyLength == 32:
		return key
	case keyLength < 16:
		return key + string(make([]byte, 16-keyLength))
	case keyLength < 24:
		return key + string(make([]byte, 24-keyLength))
	default:
		return key[:32]
	}
}
