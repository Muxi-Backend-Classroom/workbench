package service

import (
	"context"
	"github.com/Muxi-Backend-Classroom/workbench/microservice/project/dao"
	"github.com/Muxi-Backend-Classroom/workbench/pkg/errno"

	pb "github.com/Muxi-Backend-Classroom/workbench/microservice/project/proto"
)

type CreateInfo struct {
	Name      string `json:"name"`
	Intro     string `json:"intro"`
	TeamId    uint32 `json:"team_id"`
	CreatorId uint32 `json:"creator_id"`
}

func CreateProject(info *CreateInfo) error {
	project := &dao.ProjectModel{
		Name:      info.Name,
		Intro:     info.Intro,
		TeamId:    info.TeamId,
		CreatorId: info.CreatorId,
	}
	return project.Create()
}

// CreateProject ... 创建项目
func (s *ProjectService) CreateProject(ctx context.Context, req *pb.CreateProjectRequest, resp *pb.IdResponse) error {
	info := &CreateInfo{
		Name:      req.Name,
		Intro:     req.Intro,
		TeamId:    req.TeamId,
		CreatorId: req.CreatorId,
	}
	if err := CreateProject(info); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	project := &dao.ProjectModel{}
	var err error
	if project, err = dao.GetProjectByName(info.Name); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	resp.Id = uint32(project.ID)
	return nil
}
