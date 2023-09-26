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
	//使用解密并解压缩中间件
	r.Use(middlewares.DecryptAndDeCompressMiddleware())
	//该分组不校验Token
	group1 := r.Group(utils.V1Path)
	{
		//交换公钥
		group1.POST(utils.ExchangePubKeyPath, controller.ExchangePubKey)
		//共享密钥验证
		group1.POST(utils.VerifySharedKeyPath, controller.VerifySharedKey)
		//登录
		group1.POST(utils.LoginPath, controller.Login)
		//设备注册
		group1.POST(utils.DeviceRegisterPath, controller.DeviceRegister)
		//上传实时日志
		group1.POST(utils.PutOnlineLogPath, controller.PutOnlineLog)
		//查询任务
		group1.GET(utils.GetTaskPath, controller.GetTask)
		//上传日志文件
		group1.POST(utils.UploadLogFilePath, controller.UploadLogFile)
		//日志无法上传时的反馈
		group1.POST(utils.UploadLogFileErrCallBack, controller.UploadLogFileErrCallBack)
	}
	//该分组使用了校验Token的中间件
	group2 := r.Group(utils.V1Path).Use(middlewares.VerifyToken())
	{
		//创建一个应用
		group2.POST(utils.CreateAppPath, controller.CreateApp)
		//删除一个应用
		group2.POST(utils.DeleteAppPath, controller.DeleteApp)
		//查询应用列表
		group2.GET(utils.GetAppListPath, controller.GetAppList)
		//查询应用下的设备列表
		group2.GET(utils.GetDeviceListPath, controller.GetDeviceList)
		//查询设备的日志
		group2.GET(utils.GetDeviceLogPath, controller.GetDeviceLog)
		//查询任务列表
		group2.GET(utils.GetAllTaskPath, controller.GetAllTask)
		//创建任务
		group2.POST(utils.CreateTaskPath, controller.CreateTask)
		//删除任务
		group2.POST(utils.DeleteTaskPath, controller.DeleteTask)
		//查询任务的日志
		group2.GET(utils.GetTaskLogPath, controller.GetTaskLog)
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
