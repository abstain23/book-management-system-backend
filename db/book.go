package db

import (
	"book_management/dto"
	"encoding/json"
	"fmt"
	"time"

	"math/rand"
)

var bookFilePath = "storage/book.json"

func randomInt() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100000
	randomNumber := min + rand.Intn(max-min+1)
	return randomNumber
}

func GetBookList() (books []dto.Book) {
	data, err := read(bookFilePath)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &books)
	if err != nil {
		fmt.Printf("解析 book.json 数据时发生错误: %v\n", err)
	}
	return books
}

func GetBook(id int) (book dto.Book, exist bool) {
	books := GetBookList()
	for _, b := range books {
		if b.Id == id {
			book = b
			exist = true
			break
		}
	}

	return book, exist
}

func AddBook(book *dto.Book) error {
	books := GetBookList()
	book.Id = randomInt()
	books = append(books, *book)
	data, err := json.MarshalIndent(books, "", "	")
	if err != nil {
		return err
	}

	return write(bookFilePath, data)
}

func UpdateBook(book dto.Book) error {
	books := GetBookList()
	for i, b := range books {
		if b.Id == book.Id {
			books[i] = book
			break
		}
	}
	data, err := json.Marshal(books)
	if err != nil {
		return err
	}
	return write(bookFilePath, data)
}

func DeleteBook(id int) (err error, exist bool) {
	books := GetBookList()
	for i, b := range books {
		if b.Id == id {
			books = append(books[:i], books[i+1:]...)
			exist = true
			break
		}
	}
	if exist {
		data, err := json.Marshal(books)
		if err != nil {
			return err, exist
		}
		return write(bookFilePath, data), true
	}

	return nil, false
}
