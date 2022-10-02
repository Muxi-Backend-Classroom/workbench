package dao

import (
	"github.com/Muxi-Backend-Classroom/workbench/model"
	"github.com/Muxi-Backend-Classroom/workbench/pkg/constvar"
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

// Save ...
func (p *ProjectModel) Save() error {
	return model.DB.Self.Save(p).Error
}
func GetProject(id uint32) (*ProjectModel, error) {
	project := &ProjectModel{}
	if err := model.DB.Self.Table("projects").Where("id = ?", id).First(project); err != nil {
		return nil, err.Error
	}
	return project, nil
}

func GetProjectByName(name string) (*ProjectModel, error) {
	project := &ProjectModel{}
	if err := model.DB.Self.Table("projects").Where("name = ?", name).First(project); err != nil {
		return nil, err.Error
	}
	return project, nil
}

//将要删除的project的re置为1
func DeleteProject(id uint32) error {
	if err := model.DB.Self.Table("projects").Where("id = ?", id).Delete(&ProjectModel{}); err != nil {
		return err.Error
	}
	return nil
}

func ListProject(offset, limit, lastId uint32) ([]*ProjectModel, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	//list := make([]*ProjectModel, 0)
	var list []*ProjectModel
	query := model.DB.Self.Table("projects").Offset(offset).Limit(limit)

	if lastId != 0 {
		query = query.Where("id < ?", lastId).Order("id desc")
	}

	if err := query.Scan(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func ListTrashbinForDoc(offset, limit, project_id uint32) ([]*DocModel, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	var list []*DocModel
	query := model.DB.Self.Table("docs").Where("re = ? and project_id = ?", 0, project_id).Offset(offset).Limit(limit)

	if err := query.Scan(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func ListTrashbinForFile(offset, limit, project_id uint32) ([]*FileModel, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	var list []*FileModel
	query := model.DB.Self.Table("files").Where("re = ? and project_id = ?", 0, project_id).Offset(offset).Limit(limit)

	if err := query.Scan(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func ListTrashbinForMdFolder(offset, limit, project_id uint32) ([]*FolderForMdModel, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	var list []*FolderForMdModel
	query := model.DB.Self.Table("foldersformds").Where("re = ? and project_id = ?", 0, project_id).Offset(offset).Limit(limit)

	if err := query.Scan(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func ListTrashbinForFileFolder(offset, limit, project_id uint32) ([]*FolderForFileModel, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	var list []*FolderForFileModel
	query := model.DB.Self.Table("foldersforfiles").Where("re = ? and project_id = ?", 0, project_id).Offset(offset).Limit(limit)

	if err := query.Scan(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}
