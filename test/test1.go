package main

import (
	"color/utils"
	"fmt"
)

//var DB *gorm.DB
//
//func InitConfig() {
//	viper.SetConfigName("app")
//	viper.AddConfigPath("config")
//	err := viper.ReadInConfig()
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("config app.yml inited")
//}
//func InitMysql() {
//	//sql语句log打印模板
//	mylogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags),
//		logger.Config{
//			SlowThreshold: time.Second,
//			Colorful:      true,
//			LogLevel:      logger.Info,
//		},
//	)
//	var err error
//	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{
//		Logger: mylogger,
//	})
//	if err != nil {
//		fmt.Println("open fault" + err.Error())
//	}
//	DB.AutoMigrate(&models.User{})
//	DB.AutoMigrate(&models.Option{})
//	fmt.Println("mysql inited")
//}
//func main() {
//	InitConfig()
//	InitMysql()
//}

// var DB *gorm.DB
//
//	func main() {
//		viper.SetConfigName("app")
//		viper.AddConfigPath("config")
//		err := viper.ReadInConfig()
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println("config app inited")
//		newLogger := logger.New(
//			log.New(os.Stdout, "\r\n", log.LstdFlags),
//			logger.Config{
//				SlowThreshold: time.Second,
//				LogLevel:      logger.Info,
//				Colorful:      true,
//			},
//		)
//
//		DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
//			&gorm.Config{Logger: newLogger})
//		//DB.AutoMigrate(&models.User{})
//		//DB.AutoMigrate(&models.Message{})
//		DB.AutoMigrate(&models.User{})
//		DB.AutoMigrate(&models.Option{})
//	}

//	func main() {
//		password := "6550579Zsy@"
//		regex := "^(?![A-z0-9]+$)(?![A-z~@\\*()_]+$)(?![0-9~@\\*()_]+$)([A-z0-9~@\\*()_]{10,})$"
//
//		if matched, err := regexp.MatchString(regex, password); err == nil && matched {
//			fmt.Println("密码有效。")
//		} else {
//			fmt.Println("密码无效。")
//		}
//		fmt.Println("密码:", password)
//	}
func main() {
	m := utils.GenerateSMSCode()
	fmt.Println(m)
}
