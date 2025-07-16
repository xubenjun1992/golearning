package core

import (
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	DataBase string `mapstructure:"database"` // 数据库名称
	Host     string `mapstructure:"host"`     // 数据库主机地址
	Port     string `mapstructure:"port"`     // 数据库端口
	UserName string `mapstructure:"username"` // 数据库用户名
	Password string `mapstructure:"password"` // 数据库密码
	Secret   string `mapstructure:"secret"`   // JWT 密钥
	Timeout  int    `mapstructure:"timeout"`  // JWT 超时时间（秒）
}

var Configs *Config
var Once sync.Once

func init() {
	Once.Do(func() {
		Configs = &Config{}
		loadConfig()
	})
}
func loadConfig() {
	Configs = &Config{}
	viper.AddConfigPath("properties") // 配置文件所在目录
	viper.SetConfigName("config")     // 配置文件名，不带扩展名
	viper.SetConfigType("yaml")       // 配置文件类型
	if err := viper.ReadInConfig(); err != nil {
		panic("读取配置文件失败: " + err.Error())
	}
	if err := viper.Unmarshal(Configs); err != nil {
		panic("解析配置文件失败: " + err.Error())
	}
}
