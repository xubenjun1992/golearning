package core

import (
	"blog/model"
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	once sync.Once
	db   *gorm.DB
	err  error
)

func GetDb() *gorm.DB {
	once.Do(func() {
		viper.AddConfigPath("properties") // 配置文件所在目录
		viper.SetConfigName("config")     // 配置文件名，不带扩展名
		viper.SetConfigType("yaml")       // 配置文件类型
		err = viper.ReadInConfig()
		if err != nil {
			panic("读取配置文件失败: " + err.Error())
		}
		config := &Config{}
		err = viper.Unmarshal(config)
		if err != nil {
			panic("解析配置文件失败: " + err.Error())
		}
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.UserName,
			config.Password,
			config.Host,
			config.Port,
			config.DataBase,
		)

		// 建立数据库连接
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic("连接数据库失败: " + err.Error())
		}
	})
	db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
	)
	return db

}
