package service

import (
	"context"
	"github.com/Muxi-Backend-Classroom/workbench/microservice/project/dao"
	pb "github.com/Muxi-Backend-Classroom/workbench/microservice/project/proto"
	"github.com/Muxi-Backend-Classroom/workbench/pkg/errno"
)

// UpdateProjectInfo ... 更新项目信息
func (s *ProjectService) UpdateProjectInfo(ctx context.Context, req *pb.UpdateProjectInfoRequest, resp *pb.Response) error {
	project, err := dao.GetProject(req.Id)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	if project == nil {
		return errno.ServerErr(errno.ErrProjectNotExisted, err.Error())
	}

	project.Name = req.Name
	project.Intro = req.Intro

	if err := project.Save(); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
