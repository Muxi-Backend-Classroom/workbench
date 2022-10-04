package dao

import (
	"github.com/Muxi-Backend-Classroom/workbench/model"
	"github.com/jinzhu/gorm"
)

type FileModel struct {
	gorm.Model
	Url       string `gorm:"column:url;"`
	FileName  string `gorm:"column:filename;"`
	RealName  string `gorm:"column:real_name;"`
	Re        string `gorm:"column:re;"`
	Top       string `gorm:"column:top;"`
	CreatorId int    `gorm:"column:creator_id"`
	ProjectId int    `gorm:"column:project_id"`
	FatherId  int    `gorm:"column:father_id"`
}

func (FileModel) TableName() string {
	return "files"
}

// Create ...
func (f *FileModel) Create() error {
	return model.DB.Self.Create(f).Error
}
