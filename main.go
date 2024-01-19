package main

import (
	"color/router"
	"color/utils"
)

func main() {
	//InitConfig
	utils.InitConfig()
	//initMysql
	utils.InitMysql()
	//initRedis
	utils.InitRedis()
	r := router.Router()
	r.Run(":8080")
}
