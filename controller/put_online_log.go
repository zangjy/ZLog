package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func PutOnlineLog(c *gin.Context) {
	input := models.PutOnlineLogInputStruct{}
	output := models.DefaultOutputStruct{Status: utils.SuccessCode, ErrMsg: "操作成功"}
	_ = c.ShouldBindWith(&input, binding.JSON)

	sessionId := utils.GetSessionID(c)

	var tmpData []*models.OnlineLog
	for _, logStruct := range input.Data {
		tmpData = append(tmpData, &models.OnlineLog{
			SessionId:     sessionId,
			Sequence:      logStruct.Sequence,
			SystemVersion: logStruct.SystemVersion,
			AppVersion:    logStruct.AppVersion,
			TimeStamp:     logStruct.TimeStamp,
			LogLevel:      logStruct.LogLevel,
			Identify:      logStruct.Identify,
			Tag:           logStruct.Tag,
			Msg:           logStruct.Msg,
		})
	}

	if len(tmpData) > 0 {
		_ = models.WriteOnlineLogs(tmpData)
	}

	middlewares.ProcessResultData(c, output)
}
