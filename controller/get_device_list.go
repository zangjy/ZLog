package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetDeviceList(c *gin.Context) {
	input := models.GetDeviceListInputStruct{}
	output := models.GetDeviceListOutputStruct{}

	_ = c.ShouldBindWith(&input, binding.Query)

	if len(input.AppId) == 0 || input.Page <= 0 {
		output.Status = utils.ErrorCode
		output.ErrMsg = "app_id不能为空且page必须大于0"
	} else {
		output.Status = utils.SuccessCode
		output.ErrMsg = "查询成功"
		output.Data = models.GetDeviceList(input.AppId, input.Identify, input.Page)
	}

	middlewares.ProcessResultData(c, output)
}
