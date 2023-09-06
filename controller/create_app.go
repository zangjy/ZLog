package controller

import (
	"ZLog/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func CreateApp(c *gin.Context) {
	input := models.CreateAppInputStruct{}
	output := models.CreateAppOutputStruct{}
	_ = c.ShouldBindWith(&input, binding.JSON)
	if len(input.AppName) == 0 {
		output.Status = "0001"
		output.ErrMsg = "APP名称不能为空"
	} else {
		if err, appId := models.CreateApp(input.AppName); err != nil {
			output.Status = "0001"
			output.ErrMsg = "创建失败，请重试"
		} else {
			output.Status = "0000"
			output.ErrMsg = "创建成功"
			output.AppId = appId
		}
	}
	c.JSON(http.StatusOK, output)
}
