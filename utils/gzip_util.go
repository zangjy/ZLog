package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
)

// Compress 压缩输入的字节数组数据并返回压缩后的字节数组
func Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// CompressAndEncodeString 压缩输入的字符串并将其Base64编码为字符串
func CompressAndEncodeString(input string) (string, error) {
	compressedData, err := Compress([]byte(input))
	if err != nil {
		return "", err
	}
	encodedData := base64.StdEncoding.EncodeToString(compressedData)
	return encodedData, nil
}

// Decompress 解压缩输入的字节数组数据并返回解压后的字节数组
func Decompress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	buf.Write(data)
	r, err := gzip.NewReader(&buf)
	if err != nil {
		return nil, err
	}

	defer func(r *gzip.Reader) {
		_ = r.Close()
	}(r)

	var output bytes.Buffer
	_, err = io.Copy(&output, r)
	if err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}

// DecodeAndDecompressString 从Base64编码的字符串解码并解压缩数据，返回解码和解压缩后的字符串数据
func DecodeAndDecompressString(encodedData string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}

	decompressedData, err := Decompress(decodedData)
	if err != nil {
		return "", err
	}

	return string(decompressedData), nil
}
