package models

import (
	"encoding/json"
	"io"
	"strconv"
)

const (
	OK          = 0
	InvalidData = 1
	DBError     = 2
	NotFound    = 4
)

type Item struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
}

func CreateItem(r io.Reader) (int, int) {
	var item Item
	err := json.NewDecoder(r).Decode(&item)
	if err != nil {
		return InvalidData, -1
	}
	if item.Title == "" || item.Category == "" {
		return InvalidData, -1
	}
	item.ID = 0
	if err := GetDB().Create(&item).Error; err == nil {
		return DBError, -1
	}
	return OK, item.ID
}

func UpdateItem(r io.Reader, strId string) int {
	id, err := strconv.Atoi(strId)
	if err != nil {
		return InvalidData
	}
	var item Item
	err = json.NewDecoder(r).Decode(&item)
	if err != nil {
		return InvalidData
	}
	item.ID = id
	if err := GetDB().Model(&item).Updates(item).Error; err != nil {
		return DBError
	}
	return OK
}

func DeleteItem(strId string) int {
	id, err := strconv.Atoi(strId)
	if err != nil {
		return InvalidData
	}
	item := Item{ID: id}
	if err := GetDB().Delete(item).Error; err != nil {
		return DBError
	}
	return OK
}

func GetAllItems() (int, []Item) {
	var items []Item
	if err := GetDB().Find(&items).Error; err != nil {
		return DBError, nil
	}
	return OK, items
}

func GetItem(strId string) (int, Item) {
	id, err := strconv.Atoi(strId)
	if err != nil {
		return InvalidData, Item{}
	}
	var item Item
	if err := GetDB().First(&item, id).Error; err != nil {
		return NotFound, Item{}
	}
	return OK, item
}
