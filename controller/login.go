package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
		output.Status = utils.ErrorCode
		output.ErrMsg = "必要参数缺失"
	} else {
		if getUserInfoErr, user := models.GetUserInfo(input.UserName, input.Password); getUserInfoErr != nil {
			output.Status = utils.ErrorCode
			output.ErrMsg = "帐号密码不正确"
		} else {
			tokenStr, _ := newToken(user.UserName, user.Password)
			output.Status = utils.SuccessCode
			output.ErrMsg = "登录成功"
			output.Token = &tokenStr
		}
	}
	middlewares.ProcessResultData(c, output)
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
	return utils.EncryptString(tokenStr, utils.EncryptingKey)
}
