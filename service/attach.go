package service

import (
	"color/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Upload(c *gin.Context) {
	r := c.Request
	w := c.Writer
	//获得上传文件
	srcFile, head, err := r.FormFile("file")
	if err != nil {
		utils.RespFail(w, err.Error())
	}
	suffix := ".png"
	srcName := head.Filename
	t := strings.Split(srcName, ".")
	if len(t) > 1 {
		suffix = t[len(t)-1]
	}
	fileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	//创建新文件
	newFile, err := os.Create(fileName)
	if err != nil {
		utils.RespFail(w, err.Error())
	}
	//拷贝文件
	_, err := io.Copy(newFile, srcFile)

}
