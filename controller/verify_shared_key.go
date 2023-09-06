package controller

import (
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func VerifySharedKey(c *gin.Context) {
	input := models.VerifySharedKeyInputStruct{}
	output := models.VerifySharedKeyOutputStruct{}
	_ = c.ShouldBindWith(&input, binding.JSON)
	if len(input.TmpSessionID) == 0 || len(input.VerifyData) == 0 {
		output.Status = "0001"
		output.ErrMsg = "必要参数缺失"
	} else if keyPair, err := models.GetKeyPairBySessionId(input.TmpSessionID); err != nil {
		output.Status = "0001"
		output.ErrMsg = "未找到此客户端的公钥，请先和服务端进行公钥交换"
	} else if verifyData, err := utils.Decrypt(input.VerifyData, keyPair.SharedKey); err != nil {
		output.Status = "0000"
		output.ErrMsg = "验证失败"
	} else {
		output.Status = "0000"
		output.ErrMsg = "解密数据完成，请根据返回的数据进行验证"
		output.VerifyData = verifyData
	}
	c.JSON(http.StatusOK, output)
}
