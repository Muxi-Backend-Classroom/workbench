package dao

import (
	"workbench/model"
)

type StatusModel struct {
	Id      int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id" binding:"required"`
	Content string `gorm:"column:content;type:text" json:"content" binding:"required"`
	Title   string `gorm:"column:title;type:varchar(20)" json:"title" binding:"required"`
	Time    string `gorm:"column:time;type:varchar(50)" json:"time" `
	Like    int    `gorm:"column:like;type:int(11);comment:点赞数" json:"like"`
	Comment int    `gorm:"column:comment;type:int(11);comment:评论数" json:"comment"`
	UserId  int    `gorm:"column:user_id;type:int(11)" json:"user_id" binding:"required"`
}

func (StatusModel) TableName() string {
	return "status"
}

// Create ...
func (s *StatusModel) Create() error {
	return model.DB.Self.Create(s).Error
}

// Save ...
func (s *StatusModel) Save() error {
	return model.DB.Self.Save(s).Error
}
