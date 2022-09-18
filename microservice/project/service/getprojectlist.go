package service

import (
	"context"
	"github.com/Muxi-Backend-Classroom/workbench/microservice/project/dao"
	pb "github.com/Muxi-Backend-Classroom/workbench/microservice/project/proto"
	"github.com/Muxi-Backend-Classroom/workbench/pkg/errno"
)

// ...获取项目列表
func (s *ProjectService) GetProjectList(ctx context.Context, req *pb.GetProjectListRequest, resp *pb.ProjectListResponse) error {

	list, err := dao.ListProject(req.Offset, req.Limit, req.LastId)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.ProjectListItem, len(list))

	for i, item := range list {
		resList[i] = &pb.ProjectListItem{
			Id:   uint32(item.ID),
			Name: item.Name,
		}
	}

	resp.List = resList
	return nil
}
