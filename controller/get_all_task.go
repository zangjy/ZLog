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

	if len(input.AppId) == 0 || input.Page <= 0 {
		output.Status = utils.ErrorCode
		output.ErrMsg = "app_id不能为空且page必须大于0"
	} else {
		count, data := models.GetAllTask(input.AppId, input.TaskDes, input.Page)

		output.Status = utils.SuccessCode
		output.ErrMsg = "查询成功"
		output.Count = count
		output.Data = data
	}

	middlewares.ProcessResultData(c, output)
}
