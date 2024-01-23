package main

import (
	"color/models"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func main_() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app inited")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
		&gorm.Config{Logger: newLogger})
	//DB.AutoMigrate(&models.User{})
	//DB.AutoMigrate(&models.Message{})
	//DB.AutoMigrate(&models.User{})
	//DB.AutoMigrate(&models.ImageInfo{})
	//DB.AutoMigrate(&models.IssueIshida{})
	DB.AutoMigrate(&models.IssueColor{})
}
