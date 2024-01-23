package main

import (
	"color/models"
	"color/router"
	"color/utils"
	"fmt"
	"os"
	"path/filepath"
)

var fileGroups = make(map[string][]string)

func Init_color() {
	Dir := "./Asset/Upload/Color"
	err := filepath.Walk(Dir, visit)
	if err != nil {
		fmt.Println(err)
	}
	// 输出分组信息
	for group, files := range fileGroups {
		color := models.IssueColor{}
		color.Color = group
		for _, file := range files {
			if color.ImageA != "" {
				color.ImageB = "./" + file
				break
			}
			color.ImageA = "./" + file
		}
		models.AddColor(color)
	}
}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if info.IsDir() {
		return nil
	}

	// 获取文件所在的最后一级目录名
	dir := filepath.Base(filepath.Dir(path))

	// 将文件路径添加到对应的分组中
	fileGroups[dir] = append(fileGroups[dir], path)
	return nil
}

func main() {
	//InitConfig
	utils.InitConfig()
	//initMysql
	utils.InitMysql()
	//utils.DB.AutoMigrate(&models.User{})
	//utils.DB.AutoMigrate(&models.ImageInfo{})
	//utils.DB.AutoMigrate(&models.IssueIshida{})
	//utils.DB.AutoMigrate(&models.IssueColor{})
	//initRedis
	//Init_color()
	utils.InitRedis()
	r := router.Router()
	r.Run("0.0.0.0:8081")
}
