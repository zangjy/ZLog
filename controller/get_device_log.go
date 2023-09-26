package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetDeviceLog(c *gin.Context) {
	input := models.GetDeviceLogInputStruct{}
	output := models.GetDeviceLogOutputStruct{}

	_ = c.ShouldBindWith(&input, binding.Query)

	if len(input.SessionId) == 0 || input.Page <= 0 {
		output.Status = utils.ErrorCode
		output.ErrMsg = "session_id不能为空且page必须大于0"
	} else {
		output.Status = utils.SuccessCode
		output.ErrMsg = "查询成功"
		output.Data = models.GetDeviceLogs(input)
	}

	middlewares.ProcessResultData(c, output)
}
