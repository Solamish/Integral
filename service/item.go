package service

import "mobileSign/model"

type ItemInfo struct {
	Name     string `gorm:"column:name"`
	Value    string `gorm:"column:value"`
	Number   int    `gorm:"column:num"`
	PhotoSrc string `gorm:"column:photo_src"`
}

func GetItemList() (itemInfos []*ItemInfo){
	items := model.GetItemList()

	for i,_ := range items {
		itemInfo := ItemInfo{
			Name:     items[i].Name,
			Value:    items[i].Value,
			Number:   items[i].Number,
			PhotoSrc: items[i].PhotoSrc,
		}
		itemInfos = append(itemInfos, &itemInfo)
	}

	return
}

 

