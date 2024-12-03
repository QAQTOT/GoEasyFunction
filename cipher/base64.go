package cipher

import (
	"encoding/base64"
	"os"
)

// 本地图片转base64
func GetLocalImgBase64(path string) (baseImg string, err error) {
	// 获取本地文件
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	imgByte, _ := os.ReadFile(path)
	// 判断文件类型，生成一个前缀，拼接base64后可以直接粘贴到浏览器打开，不需要可以不用下面代码
	baseImg = base64.StdEncoding.EncodeToString(imgByte)
	return
}
