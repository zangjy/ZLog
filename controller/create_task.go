package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateTask(c *gin.Context) {
	input := models.CreateTaskInputStruct{}
	output := models.CreateTaskOutputStruct{}

	_ = c.ShouldBindWith(&input, binding.JSON)

	if len(input.TaskDes) == 0 || len(input.SessionId) == 0 || input.DeviceType <= 0 || input.StartTime <= 0 || input.EndTime <= 0 {
		output.Status = utils.ErrorCode
		output.ErrMsg = "task_des、session_id均不能为空且device_type、start_time、end_time必须大于0"
	} else if taskId, err := models.CreateTask(input.TaskDes, input.SessionId, input.DeviceType, input.StartTime, input.EndTime); err != nil {
		output.Status = utils.ErrorCode
		output.ErrMsg = "创建任务失败"
	} else {
		output.Status = utils.SuccessCode
		output.ErrMsg = "创建任务成功"
		output.TaskId = taskId
	}

	middlewares.ProcessResultData(c, output)
}
