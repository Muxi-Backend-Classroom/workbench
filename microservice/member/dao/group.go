package dao

import (
	m "workbench/model"
	//"workbench/pkg/constvar"

	//"github.com/jinzhu/gorm"
)

func UpdateGroupName(id uint32,name string) error {
	query := m.DB.Self.Table("groups").Where("id = ?", id)
	err := query.Update("name", name).Error
	return err
}

func UpdateGroupLeader(id uint32,leader string) error {
	query := m.DB.Self.Table("groups").Where("id = ?", id)
	err := query.Update("leader", leader).Error
	return err
}

func DeleteGroup(id uint32) error {
	query := m.DB.Self.Table("groups").Where("id = ?", id)
	err := query.Delete(&GroupModel{}).Error
	return err
}