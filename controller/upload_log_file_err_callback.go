package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func UploadLogFileErrCallBack(c *gin.Context) {
	input := models.UploadLogFileErrCallBackInputStruct{}
	output := models.DefaultOutputStruct{}
	_ = c.ShouldBindWith(&input, binding.JSON)

	sessionId := utils.GetSessionID(c)
	if len(input.TaskId) == 0 || len(input.Msg) == 0 {
		output.Status = utils.ErrorCode
		output.ErrMsg = "必要参数缺失"
	} else if state, msg := models.NotifyTaskMsg(sessionId, input.TaskId, input.Msg); !state {
		output.Status = utils.ErrorCode
		output.ErrMsg = msg
	} else {
		output.Status = utils.SuccessCode
		output.ErrMsg = "操作成功"
	}

	middlewares.ProcessResultData(c, output)
}
