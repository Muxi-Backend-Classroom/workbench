package dao

import (
	"github.com/Muxi-Backend-Classroom/workbench/model"
	"github.com/jinzhu/gorm"
)

type DocModel struct {
	gorm.Model
	FileName  string `gorm:"column:filename;"`
	Content   string `gorm:"column:content;"`
	Re        int    `gorm:"column:re;"`
	Top       int    `gorm:"column:top;"`
	EditorId  int    `gorm:"column:editor_id;"`
	CreatorId int    `gorm:"column:creator_id;"`
	ProjectId int    `gorm:"column:project_id;"`
	FatherId  int    `gorm:"column:father_id"`
}

func (DocModel) TableName() string {
	return "docs"
}

// Create ...
func (d *DocModel) Create() error {
	return model.DB.Self.Create(d).Error
}
