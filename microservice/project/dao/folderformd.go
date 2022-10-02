package dao

import "github.com/jinzhu/gorm"

type FolderForMdModel struct {
	gorm.Model
	Name      string `gorm:"column:name;"`
	Re        bool   `gorm:"column:re;"`
	CreateId  int    `gorm:"column:create_id"`
	ProjectId int    `gorm:"column:project_id"`
	FatherId  int    `gorm:"column:father_id"`
	Children  string `gorm:"column:children"`
}

func (FolderForMdModel) TableName() string {
	return "foldersformds"
}
