package quick_func

import (
	"archive/zip"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func HttpBuildQuery(params map[string]string) string {
	var queryStr []string
	for k, v := range params {
		queryStr = append(queryStr, url.QueryEscape(k)+"="+url.QueryEscape(v))
	}
	return strings.Join(queryStr, "&")
}

// DownLoadFile 下载文件
func DownLoadFile(u, savePath string) error {
	// 下载文件
	resp, err := http.Get(u)
	if err != nil {
		return errors.New("Error downloading file:" + err.Error())
	}
	defer resp.Body.Close()

	// 创建本地文件用于保存下载的ZIP文件
	outFile, err := os.Create(savePath)
	if err != nil {
		return errors.New("Error creating output file:" + err.Error())
	}
	defer outFile.Close()

	// 将下载的内容写入本地文件
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return errors.New("Error writing to output file:" + err.Error())
	}

	return nil
}

// ExtractFile 解压文件
func ExtractFile(p, out string) (error, []string) {

	r, err := zip.OpenReader(p)
	if err != nil {
		return err, nil
	}
	defer r.Close()

	var filesMap []string
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err, nil
		}
		defer rc.Close()

		fileName := f.Name
		path := filepath.Join(out, fileName)
		err = os.MkdirAll(filepath.Dir(path), f.Mode())
		if err != nil {
			return err, nil
		}
		f, err := os.OpenFile(
			path,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			f.Mode(),
		)
		if err != nil {
			return err, nil
		}
		defer f.Close()

		_, err = io.Copy(f, rc)
		if err != nil {
			return err, nil
		}

		filesMap = append(filesMap, fileName)
	}

	return nil, filesMap
}
