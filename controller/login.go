package controller

import (
	"ZLog/conf"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"time"
)

//
// Login
//  @Description: 账户登录
//  @param c
//
func Login(c *gin.Context) {
	input := models.LoginInputStruct{}
	output := models.LoginOutputStruct{}
	_ = c.ShouldBindWith(&input, binding.JSON)
	if len(input.UserName) == 0 || len(input.Password) == 0 {
		output.Status = "0001"
		output.ErrMsg = "必要参数缺失"
	} else {
		if getUserInfoErr, user := models.GetUserInfo(input.UserName, input.Password); getUserInfoErr != nil {
			output.Status = "0001"
			output.ErrMsg = "帐号密码不正确"
		} else {
			tokenStr, _ := newToken(user.UserName, user.Password)
			output.Status = "0000"
			output.ErrMsg = "登录成功"
			output.Token = &tokenStr
		}
	}
	c.JSON(http.StatusOK, output)
}

func newToken(userId, passWord string) (string, error) {
	//生效日期
	effectiveDate := time.Now().Unix()
	//Token7天后失效
	expiryDate := effectiveDate + (3600 * 24 * 7)
	//Token失效后允许在7天内凭失效Token获取一个新的Token
	periodRange := 7
	tokenStr, err := utils.NewTokenStr(userId, passWord, effectiveDate, expiryDate, periodRange)
	if err != nil {
		return "", err
	}
	return utils.Encrypt(tokenStr, conf.EncryptingKey)
}
