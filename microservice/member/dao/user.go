package dao

import (
	m "workbench/model"
	//"workbench/pkg/constvar"
	"github.com/jinzhu/gorm"
)

func GetUserRealName(id uint32) (string, error) {
	user := &UserModel{}
	d := m.DB.Self.Where("id = ?", id).First(user)
	if d.Error == gorm.ErrRecordNotFound {
		return "", nil
	}
	return user.RealName, d.Error
}
