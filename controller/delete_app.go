package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func DeleteApp(c *gin.Context) {
	input := models.DeleteAppInputStruct{}
	output := models.DefaultOutputStruct{}

	_ = c.ShouldBindWith(&input, binding.JSON)

	if len(input.AppId) == 0 {
		output.Status = utils.ErrorCode
		output.ErrMsg = "app_id不能为空"
	} else if state := models.DeleteApp(input.AppId); !state {
		output.Status = utils.ErrorCode
		output.ErrMsg = "删除失败"
	} else {
		output.Status = utils.SuccessCode
		output.ErrMsg = "删除成功"
	}

	middlewares.ProcessResultData(c, output)
}
