package service

import "mobileSign/model"

type ItemInfo struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Number   int    `json:"number"`
	PhotoSrc string `json:"photo_src"`
}

func GetItemList() (itemInfos []*ItemInfo) {
	items := model.GetItemList()

	for i, _ := range items {
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

func GetItem(name string) (*ItemInfo) {
	item := model.GetItemByName(name)
	itemInfo := &ItemInfo{
		Name:     item.Name,
		Value:    item.Value,
		Number:   item.Number,
		PhotoSrc: item.PhotoSrc,
	}
	return itemInfo
}

func DeleteItem(name string) bool {
	return model.DeleteItem(name)
}


