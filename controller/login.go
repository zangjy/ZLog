package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	if len(input.UserName) == 0 || len(input.Password) == 0 || len(input.SessionId) == 0 {
		output.Status = utils.ErrorCode
		output.ErrMsg = "必要参数缺失"
	} else {
		if getUserInfoErr, user := models.GetUserInfo(input.UserName, input.Password); getUserInfoErr != nil {
			output.Status = utils.ErrorCode
			output.ErrMsg = "帐号密码不正确"
		} else {
			tokenStr, _ := newToken(user.UserName, user.Password, input.SessionId)
			output.Status = utils.SuccessCode
			output.ErrMsg = "登录成功"
			output.Token = &tokenStr
		}
	}
	middlewares.ProcessResultData(c, output)
}

func newToken(userId, passWord, sessionId string) (string, error) {
	tokenStr, err := utils.NewTokenStr(userId, passWord, sessionId)
	if err != nil {
		return "", err
	}
	return utils.EncryptString(tokenStr, utils.EncryptingKey)
}
