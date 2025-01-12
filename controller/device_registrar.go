package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func DeviceRegister(c *gin.Context) {
	input := models.DeviceRegisterInputStruct{}
	output := models.DeviceRegisterOutputStruct{}

	_ = c.ShouldBindWith(&input, binding.JSON)

	tmpSessionId := utils.GetSessionID(c)
	if len(input.AppId) == 0 || input.DeviceType < 0 || len(input.DeviceName) == 0 || len(input.DeviceId) == 0 {
		output.Status = utils.ErrorCode
		output.ErrMsg = "app_id、device_type、device_name、device_id均不能为空"
	} else if keyPair, err := models.GetKeyPairBySessionId(tmpSessionId); err != nil {
		output.Status = utils.ErrorCode
		output.ErrMsg = "未找到此客户端的密钥对，请先和服务端进行公钥交换"
	} else if sessionId, state := models.DeviceRegister(input.AppId, input.DeviceType, input.DeviceName, input.DeviceId, keyPair.PublicKey, keyPair.SharedKey, tmpSessionId); !state {
		output.Status = utils.ErrorCode
		output.ErrMsg = "设备注册失败"
	} else {
		output.Status = utils.SuccessCode
		output.ErrMsg = "设备注册成功"
		output.SessionId = sessionId
	}

	middlewares.ProcessResultData(c, output)
}
