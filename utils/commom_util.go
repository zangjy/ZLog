package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

//
// GetSessionID
//  @Description: 从Header中取出SESSION_ID和TMP_SESSION_ID
//  @param c
//  @return string
//
func GetSessionID(c *gin.Context) string {
	s1 := c.GetHeader(SessionId)
	s2 := c.GetHeader(TmpSessionId)
	if len(s1) > 0 {
		return s1
	} else if len(s2) > 0 {
		return s2
	}
	return ""
}

//
// GetTokenFromHeader
//  @Description: 从Header里面取出Token
//  @param c
//  @return string
//
func GetTokenFromHeader(c *gin.Context) string {
	return c.GetHeader(Token)
}

// TokenStruct Token的结构体
type TokenStruct struct {
	UserId    string `json:"user_id"`
	Password  string `json:"password"`
	SessionId string `json:"session_id"`
}

func NewTokenStr(userId, pwd, sessionId string) (string, error) {
	tokenBytes, err := json.Marshal(TokenStruct{
		UserId:    userId,
		Password:  pwd,
		SessionId: sessionId,
	})
	if err != nil {
		return "", err
	}
	return string(tokenBytes), nil
}

//
// GetTokenStruct
//  @Description: 解密Token并得到Token信息的Struct
//  @param t
//  @return err
//  @return tokenStruct
//
func GetTokenStruct(t string) (err error, tokenStruct *TokenStruct) {
	tokenStruct = &TokenStruct{}
	//解密Token
	if tokenStr, err := DecryptString(t, EncryptingKey); err != nil {
		return err, nil
	} else {
		_ = json.Unmarshal([]byte(tokenStr), &tokenStruct)
		return nil, tokenStruct
	}
}

// DeleteDirectory 删除目录以及目录下的文件
func DeleteDirectory(dirPath string) error {
	//获取目录下的所有文件和子目录
	entries, err := filepath.Glob(filepath.Join(dirPath, "*"))
	if err != nil {
		return err
	}

	//删除目录下的所有文件和子目录
	for _, entry := range entries {
		err := os.Remove(entry) //调用删除文件的方法
		if err != nil {
			return err
		}
	}

	//删除目录本身
	err = os.Remove(dirPath)
	return err
}

// SaveFileToDirectory 保存文件
func SaveFileToDirectory(file *multipart.FileHeader, directory string) error {
	//创建目录（如果目录不存在）
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return err
	}

	//打开上传的文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		_ = src.Close()
	}(src)

	//创建目标文件
	dstPath := filepath.Join(directory, file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		_ = dst.Close()
	}(dst)

	//复制文件内容到目标文件
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}
