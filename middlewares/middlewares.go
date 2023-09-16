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

// DecryptMiddleware 解密中间件
//  @Description:
//  @return gin.HandlerFunc
//
func DecryptMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断是否应该进行解密
		if shouldEncryptDecrypt(c.Request.URL.Path) {
			//从Header中取出SESSION_ID和TMP_SESSION_ID
			var sessionId = getSessionID(c)
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
			//将解密后的数据写入 c.Request.Body，以供后续处理
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
	if !shouldEncryptDecrypt(c.Request.URL.Path) {
		c.JSON(http.StatusOK, data)
		return
	}
	sessionId := getSessionID(c)
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
	resultData, err := utils.EncryptString(string(jsonData), keyPair.SharedKey)
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
// shouldEncryptDecrypt
// @Description: 判断是否应该进行加密解密
// @param url
// @return bool
//
func shouldEncryptDecrypt(url string) bool {
	unNeedUrls := []string{
		utils.ExchangePubKeyPath,
		utils.VerifySharedKeyPath,
		utils.LoginPath,
		utils.CreateAppPath,
	}
	for _, unNeedUrl := range unNeedUrls {
		if strings.Contains(url, unNeedUrl) {
			return false
		}
	}
	return true
}

//
// getSessionID
//  @Description: 从Header中取出SESSION_ID和TMP_SESSION_ID
//  @param c
//  @return string
//
func getSessionID(c *gin.Context) string {
	s1 := c.GetHeader(utils.SessionId)
	s2 := c.GetHeader(utils.TmpSessionId)
	if len(s1) > 0 {
		return s1
	} else if len(s2) > 0 {
		return s2
	}
	return ""
}
