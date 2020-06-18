package main

import (
	"fmt"
	"github.com/liqing3369/poster/common"
	"github.com/liqing3369/poster/posterutil"
	"github.com/skip2/go-qrcode"
	"os"
	"time"
)

func main() {
	imageURL := "http://img.mcdsh.com/storage/images/800/O5xEIoExzNk97bRLhaIe0izqo3XbnXKi6j9BWPQb.jpeg"
	title := "title"
	content := "生成海报生成海报生成海报生成海报生成海报生成海报生成海报生成海报生成海报生成海报"
	qrURL := "https://www.baidu.com"
	posterutil.GenPosterImg(
		imageURL, title, content, qrURL, "./", "test.jpg",
	)
}
func main3() {
	genQrcode("https://www.baidu.com")
}
func main4() {
	imageURL := "http://img.mcdsh.com/storage/images/800/O5xEIoExzNk97bRLhaIe0izqo3XbnXKi6j9BWPQb.jpeg"
	title := "title"
	content := "生成海报生成海报生成海报生成海报生成海报生成海报生成海报生成海报生成海报生成海报"
	qrURL := "https://www.baidu.com"

	qrCodeURL := genQrcode(qrURL)

	// 生成海报
	style := common.Style{
		ImageURL:  imageURL,
		QRCodeURL: qrCodeURL,
		Title:     title,
		Content:   content,
	}
	postName, err := common.DrawPoster(style)
	if err != nil {
		fmt.Println("海报生成失败", err)
		os.Exit(1)
	}
	fmt.Println("运行成功")
	//fmt.Println("小程序码文件", qrcodeName)
	fmt.Println("海报文件", postName)
	os.Exit(0)
}

func genQrcode(url string) (filename string) {
	filename = fmt.Sprintf("./build/%d.png", time.Now().Local().Unix())
	err := qrcode.WriteFile(url, qrcode.Medium, 256, filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
