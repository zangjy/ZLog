package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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
	UserId   string `json:"user_id"`
	Password string `json:"password"`
	//生效日期(时间戳)
	EffectiveDate int64 `json:"effective_date"`
	//失效日期(时间戳)
	ExpiryDate int64 `json:"expiry_date"`
	//允许在一段时间范围内重新获取,天数
	PeriodRange int `json:"period_range"`
}

func NewTokenStr(userId, pwd string, effectiveDate, expiryDate int64, periodRange int) (string, error) {
	tokenBytes, err := json.Marshal(TokenStruct{
		UserId:        userId,
		Password:      pwd,
		EffectiveDate: effectiveDate,
		ExpiryDate:    expiryDate,
		PeriodRange:   periodRange,
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
	if tokenStr, err := Decrypt(t, EncryptingKey); err != nil {
		return err, nil
	} else {
		_ = json.Unmarshal([]byte(tokenStr), &tokenStruct)
		return nil, tokenStruct
	}
}
