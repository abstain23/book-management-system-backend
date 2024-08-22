package db

import (
	"book_management/dto"
	"encoding/json"
	"fmt"
)

var userFilePath = "storage/user.json"

func GetUserList() (users []dto.User) {
	data, err := read(userFilePath)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Printf("解析 user.json 数据时发生错误: %v\n", err)
	}
	return users
}

func AddUser(user dto.User) error {
	users := GetUserList()
	users = append(users, user)
	data, err := json.MarshalIndent(users, "", "	")
	if err != nil {
		fmt.Printf("序列化 JSON 数据时发生错误: %v\n", err)
		return err
	}
	return write(userFilePath, data)

}
