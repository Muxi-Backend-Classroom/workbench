package service

import (
	"context"
	p "github.com/Muxi-Backend-Classroom/workbench/microservice/project/dao"
	pb "github.com/Muxi-Backend-Classroom/workbench/microservice/project/proto"
	u "github.com/Muxi-Backend-Classroom/workbench/microservice/user/dao"
	"github.com/Muxi-Backend-Classroom/workbench/pkg/errno"
)

func (s *ProjectService) GetProjectInfo(ctx context.Context, req *pb.GetRequest, resp *pb.ProjectInfo) error {
	info, err := p.GetProject(req.Id)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	resp.Id = uint32(info.ID)
	resp.Name = info.Name
	resp.Intro = info.Intro
	resp.UserCount = info.Count
	resp.DocChildren = info.DocChildren
	resp.FileChildren = info.FileChildren
	resp.Time = info.Time
	var user *u.UserModel
	user, err = u.GetUser(info.CreatorId)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	resp.CreatorName = user.Name
	return nil
}
