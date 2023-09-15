package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"math/big"
)

// GenerateKeyPair 生成一个密钥对，包括公钥和私钥
// 返回值包括Base64编码的公钥和私钥，或者在生成密钥对过程中出现错误时返回错误
func GenerateKeyPair() (string, string, error) {
	// 使用椭圆曲线P-256生成私钥
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return "", "", err
	}

	// 将公钥编码为PKIX格式，然后进行Base64编码
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}
	encodedPublicKey := base64.StdEncoding.EncodeToString(publicKeyBytes)

	// 将私钥的D值进行Base64编码
	encodedPrivateKey := base64.StdEncoding.EncodeToString(privateKey.D.Bytes())

	return encodedPublicKey, encodedPrivateKey, nil
}

// GenerateSharedSecret 生成共享密钥
// 参数base64PublicKey和base64PrivateKey是Base64编码的公钥和私钥
// 返回值是Base64编码的共享密钥，或者在生成共享密钥过程中出现错误时返回错误
func GenerateSharedSecret(base64PublicKey, base64PrivateKey string) (string, error) {
	// 使用椭圆曲线P-256
	curve := elliptic.P256()

	// 解码Base64编码的公钥和私钥
	decodedPublicKey, err := base64.StdEncoding.DecodeString(base64PublicKey)
	if err != nil {
		return "", fmt.Errorf("base64PublicKey解码失败：%w", err)
	}
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(base64PrivateKey)
	if err != nil {
		return "", fmt.Errorf("base64PrivateKey解码失败：%w", err)
	}

	// 解析公钥
	rawPublicKey, err := x509.ParsePKIXPublicKey(decodedPublicKey)
	if err != nil {
		return "", fmt.Errorf("解析公钥失败：%w", err)
	}
	publicKey, ok := rawPublicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("无效的公钥类型")
	}

	// 创建私钥对象并生成共享密钥
	privateKey := &ecdsa.PrivateKey{
		PublicKey: *publicKey,
		D:         new(big.Int).SetBytes(decodedPrivateKey),
	}

	x, _ := curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())

	// 将共享密钥进行Base64编码
	return base64.StdEncoding.EncodeToString(x.Bytes()), nil
}
