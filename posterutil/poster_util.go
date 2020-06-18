package posterutil

import (
	"fmt"
	"github.com/liqing3369/poster/common"
	"github.com/skip2/go-qrcode"
	"os"
	"time"
)

func GenPosterImg(imageURL, title, content, qrURL, outputDIR, outputFileName string) {
	qrCodeURL := genQrcode(qrURL)
	// 生成海报
	style := common.Style{
		ImageURL:       imageURL,
		QRCodeURL:      qrCodeURL,
		Title:          title,
		Content:        content,
		OutputDIR:      outputDIR,
		OutputFileName: outputFileName,
	}
	_, err := common.DrawPoster(style)
	if err != nil {
		fmt.Println("海报生成失败", err)
		os.Exit(1)
	}
	//fmt.Printf("海报文件路径：%s\n", postName)
}

func genQrcode(url string) (filename string) {
	filename = fmt.Sprintf("./build/%d.png", time.Now().Local().Unix())
	err := qrcode.WriteFile(url, qrcode.Medium, 256, filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
