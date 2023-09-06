package conf

type AppConf struct {
	DaoConf  `ini:"dao"`
	NetConf  `ini:"net"`
	ECDHCong `ini:"ecdh"`
}

type DaoConf struct {
	UserName string `ini:"username"`
	PassWord string `ini:"password"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	DBName   string `ini:"dbname"`
}

type NetConf struct {
	Ip   string `ini:"ip"`
	Port string `ini:"port"`
}

type ECDHCong struct {
	PubKey  string `ini:"pub_key"`
	PrivKey string `ini:"priv_key"`
}

var (
	// EncryptingKey 加解密Token的密钥
	EncryptingKey = "5dfjNGwIO4Kt5C2WcS1qsApGb3c8DCyd"
	// GlobalConf 全局的配置信息
	GlobalConf = new(AppConf)
)
