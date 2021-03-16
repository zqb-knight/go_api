package mysql_model

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(maps interface{}) (tags []Tag) {
	db.Where(maps).Find(&tags)
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
