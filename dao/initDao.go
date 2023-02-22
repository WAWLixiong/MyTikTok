package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func readPassword() string {
	file, err := os.Open("../config/password")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	if err != nil {
		return ""
	}
	var password []byte
	_, err = file.Read(password)
	if err != nil {
		return ""
	}
	return string(password)
}

func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢sql阈值
			LogLevel:      logger.Error,
			Colorful:      true, // 彩色打印
		},
	)
	var err error
	password := readPassword()
	dsn := "root" + password + ":@tcp(192.168.9.128:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	// 想要正确处理time.Time 需要加上 parseTime
	// 支持emoji表请, 要使用utf8mb4
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
