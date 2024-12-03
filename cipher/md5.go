package cipher

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"reflect"
)

func GetStringMD5(s, salt string) string {
	cipher := md5.New()
	cipher.Write([]byte(s + salt))
	// 获取哈希结果
	hashBytes := cipher.Sum(nil)

	// 将哈希结果转换为十六进制字符串
	hashString := fmt.Sprintf("%x", hashBytes)

	return hashString
}

func GetFileMD5(fType string, uploadFile any) (string, int64, []byte, error) {
	// 读取文件内容
	var file *os.File
	var filePath string

	switch fType {
	case "string":
		filePath = reflect.ValueOf(uploadFile).String()
		file, err := os.Open(filePath)
		if err != nil {
			return "", 0., nil, err
		}
		defer file.Close()
		break
	case "*os.File":
		file = uploadFile.(*os.File)
		filePath = file.Name()
		break
	default:
		return "", 0, nil, errors.New("错误的文件类型")
	}
	// 计算文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		return "", 0, nil, err
	}
	fileSize := fileInfo.Size()
	// 创建一个新的MD5哈希对象
	hash := md5.New()
	// 使用io.Copy读取文件内容并同时计算哈希
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", 0, nil, err
	}
	// 获取最终的MD5哈希值
	result := hash.Sum(nil)
	// 将MD5字节切片编码为Base64字符串
	base64EncodedMD5 := base64.StdEncoding.EncodeToString(result)
	// 读取文件二进制
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", 0, nil, err
	}
	return base64EncodedMD5, fileSize, fileBytes, nil
}

func GetMultipartFileMD5(uploadFile *multipart.FileHeader) (string, int64, []byte, error) {
	// 计算文件大小
	fileSize := uploadFile.Size
	// 创建一个新的MD5哈希对象
	open, err := uploadFile.Open()
	if err != nil {
		return "", 0, nil, err
	}
	hash := md5.New()
	// 使用io.Copy读取文件内容并同时计算哈希
	_, err = io.Copy(hash, open)
	if err != nil {
		return "", 0, nil, err
	}
	// 获取最终的MD5哈希值
	result := hash.Sum(nil)
	// 将MD5字节切片编码为Base64字符串
	base64EncodedMD5 := base64.StdEncoding.EncodeToString(result)
	// 读取文件二进制
	var fileBytes []byte
	if err != nil {
		return "", 0, nil, err
	}
	defer open.Close()
	return base64EncodedMD5, fileSize, fileBytes, nil
}
