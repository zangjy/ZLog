package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetAllTask(c *gin.Context) {
	input := models.GetAllTaskInputStruct{}
	output := models.GetAllTaskOutputStruct{}

	_ = c.ShouldBindWith(&input, binding.Query)

	if input.Page <= 0 {
		output.Status = utils.ErrorCode
		output.ErrMsg = "page必须大于0"
	} else {
		output.Status = utils.SuccessCode
		output.ErrMsg = "查询成功"
		output.Data = models.GetAllTask(input.Page)
	}

	middlewares.ProcessResultData(c, output)
}
