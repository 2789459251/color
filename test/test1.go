package main

import (
	"fmt"
	"math/rand"
	"time"
)

// var fileGroups = make(map[string][]string)
//
//	func Init_color() {
//		Dir := "./Asset/Upload/Color"
//		err := filepath.Walk(Dir, visit)
//		if err != nil {
//			fmt.Println(err)
//		}
//		// 输出分组信息
//		for group, files := range fileGroups {
//			color := models.IssueColor{}
//			color.Color = group
//			for _, file := range files {
//				if color.ImageA != "" {
//					color.ImageB = "./" + file
//					break
//				}
//				color.ImageA = "./" + file
//			}
//			models.AddColor(color)
//		}
//	}
//
//	func visit(path string, info os.FileInfo, err error) error {
//		if err != nil {
//			fmt.Println(err)
//			return nil
//		}
//		if info.IsDir() {
//			return nil
//		}
//
//		// 获取文件所在的最后一级目录名
//		dir := filepath.Base(filepath.Dir(path))
//
//		// 将文件路径添加到对应的分组中
//		fileGroups[dir] = append(fileGroups[dir], path)
//		return nil
//	}
func main() {
	rand.Seed(time.Now().UnixNano())
	a := rand.Int31n(2) + 1
	fmt.Println(a)
}
