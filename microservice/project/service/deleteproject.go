package service

import (
	"context"
	"github.com/Muxi-Backend-Classroom/workbench/microservice/project/dao"
	pb "github.com/Muxi-Backend-Classroom/workbench/microservice/project/proto"
	"github.com/Muxi-Backend-Classroom/workbench/pkg/errno"
)

func (s *ProjectService) DeleteProject(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	id := req.Id
	if err := dao.DeleteProject(id); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
