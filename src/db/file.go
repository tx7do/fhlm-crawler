package db

import (
	"fmt"
	"os"

	"../g"
)

// pathExists 判断文件夹是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

// WriteToFile 写入文件
func WriteToFile(data *g.CrawlerData) {
	strFilePath := fmt.Sprintf("./db/%s/", data.Title)

	exist, err := pathExists(strFilePath)
	if err != nil {
		fmt.Printf("get dir error![%s]\n", err.Error())
		return
	}

	if !exist {
		err := os.Mkdir(strFilePath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%s]\n", err.Error())
		}
	}
}
