package main

/*添加颜色的
func main() {
	colorCode1 := "#6FA9FF"
	colorCode2 := "#657EFF"
	//创建图片
	img1 := createImage(colorCode1, 300, 300)
	img2 := createImage(colorCode2, 300, 300)

	saveImage(img1, "./Asset/Upload/Color/Blue/1.png")
	saveImage(img2, "./Asset/Upload/Color/Blue/2.png")
}
func createImage(colorCode string, width, height int) *image.RGBA {
	// 解析颜色编码
	colorRGBA, _ := hexToRGBA(colorCode)
	// 创建图像
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))
	draw.Draw(img, img.Bounds(), &image.Uniform{colorRGBA}, image.Point{}, draw.Over)
	return img
}

func saveImage(img *image.RGBA, outputPath string) {
	f, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//图片编码
	png.Encode(f, img)

}

func hexToRGBA(hexCode string) (color.RGBA, error) {
	var rgba color.RGBA
	_, err := fmt.Sscanf(hexCode, "#%02x%02x%02x", &rgba.R, &rgba.G, &rgba.B)
	if err != nil {
		return color.RGBA{}, err
	}
	rgba.A = 255 // 设置不透明度
	return rgba, nil
}

*/
