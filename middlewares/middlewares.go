package middlewares

import (
	"ZLog/models"
	"ZLog/utils"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

//
// VerifyToken
//  @Description: 校验Token的中间件
//  @return gin.HandlerFunc
//
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从Header中取出Token
		token := c.Request.Header.Get("token")

		//检查是否未传入Token
		if token == "" {
			out := models.DefaultOutputStruct{
				Status: utils.ErrorCode,
				ErrMsg: "Token不能为空",
			}
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		//Token的结构体
		var err, tokenStruct = utils.GetTokenStruct(token)
		if err != nil {
			out := models.DefaultOutputStruct{
				Status: utils.ErrorCode,
				ErrMsg: "无效Token",
			}
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		//校验Token是否已过期
		_, state := utils.Get(tokenStruct.SessionId)
		if !state {
			out := models.DefaultOutputStruct{
				Status: utils.ErrorCode,
				ErrMsg: "Token已过期",
			}
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		//查询该账户信息
		if err, _ := models.GetUserInfo(tokenStruct.UserId, tokenStruct.Password); err != nil {
			out := models.DefaultOutputStruct{
				Status: utils.ErrorCode,
				ErrMsg: "帐号不存在",
			}
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}
	}
}

// DecryptAndDeCompressMiddleware 解密并解压缩中间件
//  @Description:
//  @return gin.HandlerFunc
//
func DecryptAndDeCompressMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断是否应该进行解密并解压缩
		if shouldProcessData(c.Request.URL.Path) && c.Request.Method != "GET" {
			//从Header中取出SESSION_ID和TMP_SESSION_ID
			var sessionId = utils.GetSessionID(c)
			//如果SESSION_ID和TMP_SESSION_ID都为空，则返回错误
			if len(sessionId) == 0 {
				ProcessResultData(c, models.DefaultOutputStruct{
					Status: utils.ErrorCode,
					ErrMsg: "SESSION_ID不能为空",
				})
				c.Abort()
				return
			}
			//从缓存中取出密钥对
			keyPair, err := models.GetKeyPairBySessionId(sessionId)
			if err != nil {
				ProcessResultData(c, models.DefaultOutputStruct{
					Status: utils.ErrorCode,
					ErrMsg: "未找到此客户端的密钥对，请先和服务端进行公钥交换",
				})
				c.Abort()
				return
			}
			//读取请求体数据
			body, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				ProcessResultData(c, models.DefaultOutputStruct{
					Status: utils.ErrorCode,
					ErrMsg: "无法读取请求体数据",
				})
				c.Abort()
				return
			}
			//解密请求体数据
			resultData, err := utils.DecryptString(string(body), keyPair.SharedKey)
			if err != nil {
				ProcessResultData(c, models.DefaultOutputStruct{
					Status: utils.ErrorCode,
					ErrMsg: "数据解密失败",
				})
				c.Abort()
				return
			}
			//解压缩数据
			resultData, err = utils.DecompressString(resultData)
			if err != nil {
				ProcessResultData(c, models.DefaultOutputStruct{
					Status: utils.ErrorCode,
					ErrMsg: "数据解压缩失败",
				})
				c.Abort()
				return
			}
			//将解密并解压缩后的数据写入 c.Request.Body，以供后续处理
			c.Request.Body = ioutil.NopCloser(bytes.NewReader([]byte(resultData)))
		}
		//继续处理
		c.Next()
	}
}

//
// ProcessResultData 处理响应数据
//  @Description:
//  @param c
//  @param data
//
func ProcessResultData(c *gin.Context, data interface{}) {
	if !shouldProcessData(c.Request.URL.Path) {
		c.JSON(http.StatusOK, data)
		return
	}
	sessionId := utils.GetSessionID(c)
	if len(sessionId) == 0 {
		c.JSON(http.StatusOK, models.DefaultOutputStruct{
			Status: utils.ErrorCode,
			ErrMsg: "SESSION_ID不能为空",
		})
		return
	}
	keyPair, err := models.GetKeyPairBySessionId(sessionId)
	if err != nil {
		c.JSON(http.StatusOK, models.DefaultOutputStruct{
			Status: utils.ErrorCode,
			ErrMsg: "未找到此客户端的密钥对，请先和服务端进行公钥交换",
		})
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusOK, models.DefaultOutputStruct{
			Status: utils.ErrorCode,
			ErrMsg: "无法序列化响应数据",
		})
		return
	}
	resultData, err := utils.CompressString(string(jsonData))
	if err != nil {
		c.JSON(http.StatusOK, models.DefaultOutputStruct{
			Status: utils.ErrorCode,
			ErrMsg: "数据压缩失败",
		})
		return
	}
	resultData, err = utils.EncryptString(resultData, keyPair.SharedKey)
	if err != nil {
		c.JSON(http.StatusOK, models.DefaultOutputStruct{
			Status: utils.ErrorCode,
			ErrMsg: "数据加密失败",
		})
		return
	}
	c.String(http.StatusOK, resultData)
}

//
// shouldProcessData
// @Description: 根据URL判断是否需要加密、解密、压缩、解压缩数据
// @param url
// @return bool
//
func shouldProcessData(url string) bool {
	unNeedUrls := []string{
		utils.ExchangePubKeyPath,
		utils.VerifySharedKeyPath,
		utils.UploadLogFilePath,
	}
	for _, unNeedUrl := range unNeedUrls {
		if strings.HasSuffix(url, unNeedUrl) {
			return false
		}
	}
	return true
}
