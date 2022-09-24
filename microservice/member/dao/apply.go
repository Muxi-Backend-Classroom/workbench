package dao

import (
	m "workbench/model"
	//"workbench/pkg/constvar"
	//"github.com/jinzhu/gorm"
)

func DeleteApply(id uint32) error {
	query := m.DB.Self.Table("applys").Where("id = ?", id)
	err := query.Delete(&ApplyModel{}).Error
	return err
}
