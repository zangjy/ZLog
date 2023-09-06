package controller

import (
	"ZLog/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func DeviceRegister(c *gin.Context) {
	input := models.DeviceRegisterInputStruct{}
	output := models.DeviceRegisterOutputStruct{}
	_ = c.ShouldBindWith(&input, binding.JSON)
	if len(input.AppId) == 0 || input.DeviceType < 0 || len(input.DeviceName) == 0 || len(input.DeviceId) == 0 || len(input.TmpSessionID) == 0 {
		output.Status = "0001"
		output.ErrMsg = "必要参数缺失"
	} else if keyPair, err := models.GetKeyPairBySessionId(input.TmpSessionID); err != nil {
		output.Status = "0001"
		output.ErrMsg = "未找到此客户端的公钥，请先和服务端进行公钥交换"
	} else if sessionId, state := models.DeviceRegister(input.AppId, input.DeviceType, input.DeviceName, input.DeviceId, keyPair.PublicKey, keyPair.SharedKey, input.TmpSessionID); !state {
		output.Status = "0001"
		output.ErrMsg = "设备注册失败"
	} else {
		output.Status = "0000"
		output.ErrMsg = "设备注册成功"
		output.SessionId = sessionId
	}
	c.JSON(http.StatusOK, output)
}
