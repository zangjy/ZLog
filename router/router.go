package router

import (
	"ZLog/controller"
	"ZLog/middlewares"
	"ZLog/models"
	"ZLog/utils"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"net/http"
)

//
// SetUpRouter
//  @Description: 初始化路由
//  @param addr 监听的地址
//  @return err err不为nil是代表初始化失败
//
func SetUpRouter(addr string) (err error) {
	//设置Gin的模式
	gin.SetMode(gin.ReleaseMode)
	//解决乱码问题
	gin.DefaultWriter = colorable.NewColorableStdout()
	//日志带颜色
	gin.ForceConsoleColor()
	//默认配置
	r := gin.Default()
	//使用解密中间件
	r.Use(middlewares.DecryptMiddleware())
	//该分组不校验Token
	group1 := r.Group(utils.V1Path)
	{
		//登录
		group1.POST(utils.LoginPath, controller.Login)
		//交换公钥
		group1.POST(utils.ExchangePubKeyPath, controller.ExchangePubKey)
		//共享密钥验证
		group1.POST(utils.VerifySharedKeyPath, controller.VerifySharedKey)
		//设备注册
		group1.POST(utils.DeviceRegisterPath, controller.DeviceRegister)
	}
	//该分组使用了校验Token的中间件
	group2 := r.Group(utils.V1Path).Use(middlewares.Verify())
	{
		//创建一个应用
		group2.POST(utils.CreateAppPath, controller.CreateApp)
	}
	//未匹配到路由时的处理
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, models.DefaultOutputStruct{
			Status: utils.StatusNotFoundCode,
			ErrMsg: "未找到该路由",
		})
	})
	err = r.Run(addr)
	return
}
