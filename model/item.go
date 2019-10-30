package model

import (
	"fmt"
	"log"
)

type Item struct {
	BaseModel
	Name     string `gorm:"column:name"`
	Value    string `gorm:"column:value"`
	Number   int    `gorm:"column:num"`
	PhotoSrc string `gorm:"column:photo_src"`
}

type Photo struct {
	IsUpload bool
	PhotoSrc string
}

var photoChan chan Photo

// 获得全部商品信息
func GetItemList() (items []*Item) {
	rows, err := DB.Raw("select * from items").Rows()
	if err != nil {
		log.Println("fail to select items", err)
	}
	defer rows.Close()
	for rows.Next() {
		var item Item
		scanErr := DB.ScanRows(rows, &item)
		if scanErr != nil {
			log.Println("fail to scan item to struct", scanErr)
		}
		items = append(items, &item)
	}
	return
}

// 添加
func (item *Item) AddItem() bool {
	tmp := Item{}
	DB.Where(" name = ?", item.Name).Find(&tmp)

	// 同名商品不能上传
	if tmp.Name != "" {
		return false
	}

	// 价格和数量不能为空
	if item.Value == "" || item.Number == 0 {
		return false
	}

	// 上传
	err := DB.Create(item).Error
	if err != nil {
		log.Println("fail to add item", err)
		return false
	} else {
		go func() {
			p := Photo{
				IsUpload: true,
				PhotoSrc: item.PhotoSrc,
			}
			photoChan <- p
		}()
	}
	return true
}

// 删除
func DeleteItem(name string) bool {
	err := DB.Where("name = ?", name).Delete(&Item{}).Error
	if err != nil {
		log.Println("fail to delete item", err)
		return false
	}
	return true
}

// 更新商品
func (item *Item) UpdateItems() bool {

	tmp := Item{}
	DB.Where("select * from items where name = ?", item.Name).Find(&tmp)

	// 商品不存在，返回false
	if tmp.Name == "" {
		return false
	}

	err := DB.Table("items").Updates(item).Error
	if err != nil {
		fmt.Println("fail to update item", err)
		return false
	}
	return true
}

// 获得单个商品的信息
func GetItemByName(name string) (Item) {
	var item Item
	err := DB.Raw("select * from items where name = ?", name).Scan(&item).Error
	if err != nil {
		fmt.Println("fail to get item info", err)
	}
	return item
}

// 上传图片
func Upload() {
	go func() {
		for {
			<-photoChan
			// TODO 上传图片
		}
	}()
}
