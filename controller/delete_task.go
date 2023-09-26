package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func DeleteTask(c *gin.Context) {
	input := models.DeleteTaskInputStruct{}
	output := models.DefaultOutputStruct{}

	_ = c.ShouldBindWith(&input, binding.JSON)

	if len(input.TaskId) == 0 {
		output.Status = utils.ErrorCode
		output.ErrMsg = "task_id不能为空"
	} else if err := models.DeleteTask(input.TaskId); err != nil {
		output.Status = utils.ErrorCode
		output.ErrMsg = "删除任务失败"
	} else {
		output.Status = utils.SuccessCode
		output.ErrMsg = "删除任务成功"
	}

	middlewares.ProcessResultData(c, output)
}
