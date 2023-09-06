package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
)

// CompressAndEncodeString 压缩并对输入字符串进行Base64编码。
// 返回值是压缩和编码后的字符串，或者在发生错误时返回一个错误。
func CompressAndEncodeString(input string) (string, error) {
	// 创建一个缓冲区用于存储压缩后的数据
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	// 将输入字符串写入压缩器
	_, err := gz.Write([]byte(input))
	if err != nil {
		return "", err
	}

	// 关闭压缩器以确保所有数据都被写入缓冲区
	if err := gz.Close(); err != nil {
		return "", err
	}

	// 对压缩后的数据进行Base64编码
	encodedData := base64.StdEncoding.EncodeToString(buf.Bytes())
	return encodedData, nil
}

// DecodeAndDecompressString 解码Base64字符串并解压缩它。
// 返回值是解压缩后的字符串，或者在发生错误时返回一个错误。
func DecodeAndDecompressString(encodedData string) (string, error) {
	// 解码Base64字符串
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}

	// 创建一个gzip.Reader来解压缩数据
	r, err := gzip.NewReader(bytes.NewReader(decodedData))
	if err != nil {
		return "", err
	}
	defer func(r *gzip.Reader) {
		_ = r.Close()
	}(r)

	// 将解压缩后的数据复制到缓冲区
	var decompressedData bytes.Buffer
	_, err = io.Copy(&decompressedData, r)
	if err != nil {
		return "", err
	}

	// 将解压缩后的数据转换为字符串并返回
	return decompressedData.String(), nil
}
