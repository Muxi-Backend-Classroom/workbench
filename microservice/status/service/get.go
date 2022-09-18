package service

import (
	"context"

	"github.com/Muxi-Backend-Classroom/workbench/microservice/status/dao"
	pb "github.com/Muxi-Backend-Classroom/workbench/microservice/status/proto"
	"github.com/Muxi-Backend-Classroom/workbench/pkg/errno"
)

// Get ... 获取单个的进度信息
func (s *StatusService) Login(ctx context.Context, req *pb.GetRequest, resp *pb.Status) error {
	status, err := dao.GetStatus(req.Id, req.Uid)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	if status == nil {
		return errno.ServerErr(errno.ErrstatusNotExisted, "")
	}
	resp.Id = uint32(status.ID)
	resp.Title = status.Title
	resp.Content = status.Content
	resp.UserId = status.User_id
	resp.Time = status.Time
	resp.Avatar = status.Avatar
	resp.UserName = status.User_name
	resp.Liked = status.Liked
	resp.Comment = status.Comment
	resp.Like = status.Like

	return nil
}
