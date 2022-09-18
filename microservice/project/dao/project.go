package dao

import (
	"github.com/Muxi-Backend-Classroom/workbench/model"
	"github.com/jinzhu/gorm"
)

type ProjectModel struct {
	gorm.Model
	Name         string `gorm:"column:name;"`
	Intro        string `gorm:"column:intro;"`
	Time         string `gorm:"column:time;"`
	Count        uint32 `gorm:"column:count;"`
	TeamId       uint32 `gorm:"column:team_id;"`
	FileChildren string `gorm:"column:file_children;"`
	DocChildren  string `gorm:"column:doc_children;"`
	CreatorId    uint32 `gorm:"column:creator_id;"`
}

func (ProjectModel) TableName() string {
	return "projects"
}

// Create ...
func (p *ProjectModel) Create() error {
	return model.DB.Self.Create(p).Error
}

func GetProject(id uint32) (*ProjectModel, error) {
	project := &ProjectModel{}
	if err := model.DB.Self.Where("id = ?", id).First(project); err != nil {
		return nil, err.Error
	}
	return project, nil
}
