package db

import (
	"fmt"
	"os"
)

func init() {
	initFile(userFilePath)
	initFile(bookFilePath)
}

func read(path string) (data []byte, err error) {
	data, err = os.ReadFile(path)
	if err != nil {
		fmt.Printf("读取文件 %s 时发生错误: %v\n", path, err)
	}
	return data, err
}

func write(path string, data []byte) error {
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Printf("写入文件 %s 时发生错误: %v\n", path, err)
	}
	return err
}

func initFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.WriteFile(path, []byte("[]"), 0644)
		if err != nil {
			fmt.Printf("创建文件 %s 时发生错误: %v\n", path, err)
			return
		}
		fmt.Printf("文件 %s 已创建。\n", path)
	} else if err != nil {
		fmt.Printf("检查文件 %s 时发生错误: %v\n", path, err)
	} else {
		fmt.Printf("文件 %s 已存在。\n", path)
	}
}
