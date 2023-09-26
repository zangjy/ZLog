package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetTask(c *gin.Context) {
	input := models.GetTaskInputStruct{}
	output := models.GetTaskOutputStruct{}

	_ = c.ShouldBindWith(&input, binding.Query)

	sessionId := utils.GetSessionID(c)

	output.Status = utils.SuccessCode
	output.ErrMsg = "查询成功"
	output.Data = models.GetTaskList(sessionId, input.DeviceType)

	middlewares.ProcessResultData(c, output)
}
