package router

import (
	"ZLog/controller"
	"ZLog/middlewares"
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
	gin.SetMode(gin.DebugMode)
	//解决乱码问题
	gin.DefaultWriter = colorable.NewColorableStdout()
	//日志带颜色
	gin.ForceConsoleColor()
	//默认配置
	r := gin.Default()
	//该分组不校验Token
	group1 := r.Group(V1Path)
	{
		//登录
		group1.POST(LoginPath, controller.Login)
		//交换公钥
		group1.POST(ExchangePubKeyPath, controller.ExchangePubKey)
		//共享密钥验证
		group1.POST(VerifySharedKeyPath, controller.VerifySharedKey)
		//绑定设备
		group1.POST(DeviceRegisterPath, controller.DeviceRegister)
	}
	//该分组使用了校验Token的中间件
	group2 := r.Group(V1Path).Use(middlewares.Verify())
	{
		//创建一个应用
		group2.POST(CreateAppPath, controller.CreateApp)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"err_msg": "无效路径",
			"status":  "404",
		})
	})
	err = r.Run(addr)
	return
}
