package main

import (
	"ZLog/conf"
	"ZLog/dao"
	"ZLog/models"
	"ZLog/router"
	"ZLog/utils"
	"fmt"
	"gopkg.in/ini.v1"
)

func main() {
	fmt.Println("********************")
	fmt.Println("开发者:ZangJiaYu")
	fmt.Println("系统版本:1.0")
	fmt.Println("********************")
	
	// 获取或创建ECDH配置节
	if iniFile, loadFileErr := ini.Load("./conf/conf.ini"); loadFileErr != nil {
		fmt.Printf("加载配置文件失败:%v\n", loadFileErr)
		return
	} else {
		section, getSectionErr := iniFile.GetSection("ecdh")
		if getSectionErr != nil {
			// 如果配置节不存在，则创建它
			section, _ = iniFile.NewSection("ecdh")
		}
		//获取"pub_key"和"priv_key"的值
		pubKey := section.Key("pub_key").String()
		privKey := section.Key("priv_key").String()
		//如果"pub_key"或"priv_key"为空，则覆写它们
		if pubKey == "" || privKey == "" {
			if pubKey, privKey, generateKeyPairErr := utils.GenerateKeyPair(); generateKeyPairErr != nil {
				fmt.Printf("生成密钥对失败：%v", generateKeyPairErr)
				return
			} else {
				if _, pubKeyErr := section.NewKey("pub_key", pubKey); pubKeyErr != nil {
					fmt.Printf("创建公钥失败：%v", generateKeyPairErr)
					return
				}
				if _, privKeyErr := section.NewKey("priv_key", privKey); privKeyErr != nil {
					fmt.Printf("创建私钥失败：%v", generateKeyPairErr)
					return
				}
			}
		}
		//保存配置文件
		if saveErr := iniFile.SaveTo("./conf/conf.ini"); saveErr != nil {
			fmt.Printf("保存配置文件失败：%v", saveErr)
			return
		}
	}

	//读取配置文件
	if err := ini.MapTo(conf.GlobalConf, "./conf/conf.ini"); err != nil {
		fmt.Printf("读取配置文件失败:%v\n", err)
		return
	}

	//连接数据库
	if err := dao.InitDao(conf.GlobalConf.DaoConf.UserName, conf.GlobalConf.DaoConf.PassWord, conf.GlobalConf.DaoConf.Host, conf.GlobalConf.DaoConf.Port, conf.GlobalConf.DaoConf.DBName); err != nil {
		fmt.Printf("数据库连接失败:%v\n", err)
		return
	}

	//数据库表初始化
	if err := dao.DB.AutoMigrate(&models.User{}, &models.App{}, &models.Device{}, &models.OnlineLog{}, &models.Task{}, &models.OfflineLog{}); err != nil {
		fmt.Printf("数据库模型创建失败:%v\n", err)
		return
	}

	//初始化Gin
	if err := router.SetUpRouter(conf.GlobalConf.NetConf.Ip + ":" + conf.GlobalConf.NetConf.Port); err != nil {
		fmt.Printf("启动服务失败:%v\n", err)
		return
	}
}
