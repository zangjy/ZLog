package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

//
// GetTokenFromHeader
//  @Description: 从Header里面取出Token
//  @param c
//  @return string
//
func GetTokenFromHeader(c *gin.Context) string {
	return c.Request.Header.Get("token")
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
