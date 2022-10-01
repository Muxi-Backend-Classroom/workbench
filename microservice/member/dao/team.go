package dao

import (
	m "workbench/model"
	//"workbench/pkg/constvar"

	//"github.com/jinzhu/gorm"
)

func UpdateTeamName(id uint32,name string) error {
	query := m.DB.Self.Table("teams").Where("id = ?", id)
	err := query.Update("name", name).Error
	return err
}

func DeleteTeam(id uint32) error {
	query := m.DB.Self.Table("teams").Where("id = ?", id)
	err := query.Delete(&TeamModel{}).Error
	return err
}