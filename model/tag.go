package model

import "log"

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(maps interface{}) (tags []Tag, err error) {
	err = db.Where(maps).Find(&tags).Error
	if err != nil {
		log.Println("数据库查询出错")
	}
	log.Println(tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func AddTag(name string, createdBy string) error {
	tag := Tag{
		Name:      name,
		CreatedBy: createdBy,
	}
	err := db.Create(&tag).Error
	return err
}

func UpdateTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}
