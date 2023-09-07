package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateApp(c *gin.Context) {
	input := models.CreateAppInputStruct{}
	output := models.CreateAppOutputStruct{}
	_ = c.ShouldBindWith(&input, binding.JSON)
	if len(input.AppName) == 0 {
		output.Status = utils.ErrorCode
		output.ErrMsg = "APP名称不能为空"
	} else {
		if err, appId := models.CreateApp(input.AppName); err != nil {
			output.Status = utils.ErrorCode
			output.ErrMsg = "创建失败，请重试"
		} else {
			output.Status = utils.SuccessCode
			output.ErrMsg = "创建成功"
			output.AppId = appId
		}
	}
	middlewares.ProcessResultData(c, output)
}
