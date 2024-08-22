package db

import (
	"book_management/dto"
	"encoding/json"
	"fmt"
	"os"
)

var filePath = "storage/user.json"

func init() {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("创建文件 %s 时发生错误: %v\n", filePath, err)
			return
		}
		defer file.Close()
		fmt.Printf("文件 %s 已创建。\n", filePath)
	} else if err != nil {
		fmt.Printf("检查文件 %s 时发生错误: %v\n", filePath, err)
	} else {
		fmt.Printf("文件 %s 已存在。\n", filePath)
	}
}

func GetUserList() (users []dto.User) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("读取文件 %s 时发生错误: %v\n", filePath, err)
		return
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Printf("解析 JSON 数据时发生错误: %v\n", err)
		return
	}

	return users
}

func AddUser(user dto.User) error {
	users := GetUserList()
	users = append(users, user)

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Printf("序列化 JSON 数据时发生错误: %v\n", err)
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		fmt.Printf("写入文件 %s 时发生错误: %v\n", filePath, err)
		return err
	}

	return nil
}
