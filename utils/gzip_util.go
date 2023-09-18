package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"errors"
	"io"
)

// CompressBytes 压缩字节数组
func CompressBytes(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}

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

// CompressString 压缩字符串并进行Base64编码
func CompressString(data string) (string, error) {
	if len(data) == 0 {
		return "", errors.New("data is empty")
	}

	compressedData, err := CompressBytes([]byte(data))
	if err != nil {
		return "", err
	}
	encodedData := base64.StdEncoding.EncodeToString(compressedData)
	return encodedData, nil
}

// DecompressBytes 解压缩字节数组数组
func DecompressBytes(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}

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

// DecompressString Base64解码数据并进行解压缩数据
func DecompressString(data string) (string, error) {
	if len(data) == 0 {
		return "", errors.New("data is empty")
	}

	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	decompressedData, err := DecompressBytes(decodedData)
	if err != nil {
		return "", err
	}

	return string(decompressedData), nil
}
