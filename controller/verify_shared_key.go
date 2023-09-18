package controller

import (
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func VerifySharedKey(c *gin.Context) {
	input := models.VerifySharedKeyInputStruct{}
	output := models.VerifySharedKeyOutputStruct{}
	_ = c.ShouldBindWith(&input, binding.JSON)

	if len(input.TmpSessionId) == 0 || len(input.VerifyData) == 0 {
		output.Status = utils.ErrorCode
		output.ErrMsg = "必要参数缺失"
	} else if keyPair, err := models.GetKeyPairBySessionId(input.TmpSessionId); err != nil {
		output.Status = utils.ErrorCode
		output.ErrMsg = "未找到此客户端的密钥对，请先和服务端进行公钥交换"
	} else if decryptData, err := utils.DecryptString(input.VerifyData, keyPair.SharedKey); err != nil {
		output.Status = utils.SuccessCode
		output.ErrMsg = "验证失败"
	} else {
		output.Status = utils.SuccessCode
		output.ErrMsg = "服务端已对数据进行解密，如果解密结果正确，则说明共享密钥验证成功，否则请检查客户端代码并再次尝试"
		output.DecryptData = decryptData
	}

	middlewares.ProcessResultData(c, output)
}
