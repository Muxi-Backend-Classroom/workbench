package service

import (
	"context"
	//"workbench/pkg/errno"

	//"workbench-team/pkg"
	"time"
	"workbench-team/dao"
	pb "workbench-team/proto"
)

// NewGroup ... 创建
func (s *MemberService) NewGroup(ctx context.Context, req *pb.NewGroupRequest, resp *pb.Response) error {

	name, err := dao.GetUserRealName(req.UserId)
	if err != nil {
		return err
	}
	group := dao.GroupModel{
		Name:    req.Name,
		Count:   1,
		Time:    time.Now().String(),
		Leader:  name,
	}
	group.Save()

	return nil
}

// UpdateGroup ... 更新
func (s *MemberService) UpdateGroup(ctx context.Context, req *pb.UpdateGroupRequest, resp *pb.Response) error {

	err := dao.UpdateGroupName(req.Id, req.Name)
	if err != nil {
		return err
	}

	return nil
}

// DeleteGroup ... 删除
func (s *MemberService) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest, resp *pb.Response) error {

	err := dao.DeleteGroup(req.Id)
	if err != nil {
		return err
	}

	return nil
}