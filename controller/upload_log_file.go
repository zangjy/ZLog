package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func UploadLogFile(c *gin.Context) {
	output := models.DefaultOutputStruct{}

	file, err := c.FormFile("file")
	if err != nil {
		output.Status = utils.ErrorCode
		output.ErrMsg = "接收文件失败"
	} else {
		sessionId := utils.GetSessionID(c)

		//获取文件名
		fileName := file.Filename
		//文件路径
		filePath := utils.LogFileRootPath + "/" + fileName
		//不包含后缀的文件名作为taskId
		taskId := filepath.Base(fileName[:len(fileName)-len(filepath.Ext(fileName))])

		if err := utils.SaveFileToDirectory(file, utils.LogFileRootPath); err != nil {
			output.Status = utils.ErrorCode
			output.ErrMsg = "保存文件失败"
		} else if notifyState, msg := models.NotifyTaskState(sessionId, taskId, 1); !notifyState {
			output.Status = utils.ErrorCode
			output.ErrMsg = msg
			//删除文件
			_ = os.Remove(filePath)
		} else {
			output.Status = utils.SuccessCode
			output.ErrMsg = "操作成功"
			//添加到日志解析任务队列
			ZLogProcessorInstance.AddTask(filePath)
		}
	}

	middlewares.ProcessResultData(c, output)
}
