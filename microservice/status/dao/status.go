package dao

import (
	"github.com/Muxi-Backend-Classroom/workbench/model"
	"github.com/Muxi-Backend-Classroom/workbench/pkg/constvar"
	"github.com/jinzhu/gorm"
)

type StatusModel struct {
	gorm.Model
	Title     string `gorm:"column:title;"`
	Content   string `gorm:"column:content;"`
	User_id   uint32 `gorm:"column:user_id;"`
	Time      uint32 `gorm:"column:time;"`
	Avatar    string `gorm:"column:avatar;"`
	User_name string `gorm:"column:user_name;"`
	Liked     bool   `gorm:"column:liked;"`
	Comment   uint32 `gorm:"column:comment;"`
	Like      uint32 `gorm:"column:like;"`
}

func (StatusModel) TableName() string {
	return "statuss"
}

// Create ...
func (s *StatusModel) Create() error {
	return model.DB.Self.Create(s).Error
}

// Save ...
func (s *StatusModel) Save() error {
	return model.DB.Self.Save(s).Error
}
func GetStatus(id uint32, uid uint32) (*StatusModel, error) {
	Status := &StatusModel{}
	if err := model.DB.Self.Table("Status").Where("id = ? and user_id = ?", id, uid).First(Status); err != nil {
		return nil, err.Error
	}
	return Status, nil
}

func GetStatusByName(name string) (*StatusModel, error) {
	Status := &StatusModel{}
	if err := model.DB.Self.Table("Status").Where("name = ?", name).First(Status); err != nil {
		return nil, err.Error
	}
	return Status, nil
}

func DeleteStatus(id uint32) error {
	if err := model.DB.Self.Table("Status").Where("id = ?", id).Delete(&StatusModel{}); err != nil {
		return err.Error
	}
	return nil
}

func ListStatus(offset, limit, lastId uint32) ([]*StatusModel, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	list := make([]*StatusModel, 0)

	query := model.DB.Self.Table("Status").Offset(offset).Limit(limit)

	if lastId != 0 {
		query = query.Where("id < ?", lastId).Order("id desc")
	}

	if err := query.Scan(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}
