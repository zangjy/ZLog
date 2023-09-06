package middlewares

import (
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//
// Verify
//  @Description: 路由中间件，校验Token，通过则放行
//  @return gin.HandlerFunc
//
func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从Header中取出Token
		token := c.Request.Header.Get("token")

		//检查是否未传入Token
		if token == "" {
			out := models.DefaultOutputStruct{
				Status: "0001",
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
				Status: "0002",
				ErrMsg: "无效Token",
			}
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		//校验Token是否已过期
		if (time.Now().Unix() - tokenStruct.ExpiryDate) > 0 {
			out := models.DefaultOutputStruct{
				Status: "0002",
				ErrMsg: "Token已过期",
			}
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}

		//查询该账户信息
		if err, _ := models.GetUserInfo(tokenStruct.UserId, tokenStruct.Password); err != nil {
			out := models.DefaultOutputStruct{
				Status: "0001",
				ErrMsg: "帐号不存在",
			}
			c.JSON(http.StatusOK, out)
			c.Abort()
			return
		}
	}
}
