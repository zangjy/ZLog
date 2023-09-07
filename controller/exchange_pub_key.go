package controller

import (
	"ZLog/conf"
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ExchangePubKey(c *gin.Context) {
	input := models.ExchangePubKeyInputStruct{}
	output := models.ExchangePubKeyOutputStruct{}
	_ = c.ShouldBindWith(&input, binding.JSON)
	if len(input.ClientPubKey) == 0 {
		output.Status = utils.ErrorCode
		output.ErrMsg = "必要参数缺失"
	} else if sharedKey, err := utils.GenerateSharedSecret(input.ClientPubKey, conf.GlobalConf.ECDHCong.PrivKey); err != nil {
		output.Status = utils.ErrorCode
		output.ErrMsg = "未成功生成共享密钥，请重试"
	} else {
		//生成sessionId
		sessionId := utils.WorkerInstance.GetId()
		//将临时的SessionId对应的客户端公钥和共享密钥存入Map中
		utils.Put(sessionId, input.ClientPubKey, sharedKey)
		//返回数据
		output.Status = utils.SuccessCode
		output.ErrMsg = "操作成功"
		output.TmpSessionID = sessionId
		output.ServerPubKey = conf.GlobalConf.ECDHCong.PubKey
	}
	middlewares.ProcessResultData(c, output)
}
