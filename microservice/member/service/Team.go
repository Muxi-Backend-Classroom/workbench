package service

import (
	"context"
	//"workbench/pkg/errno"

	//"workbench-team/pkg"
	"time"
	"workbench-team/dao"
	pb "workbench-team/proto"
)

// NewTeam ... 创建新团队
func (s *MemberService) NewTeam(ctx context.Context, req *pb.NewTeamRequest, resp *pb.Response) error {

	name, err := dao.GetUserRealName(req.UserId)
	if err != nil {
		return err
	}
	team := dao.TeamModel{
		Name:    req.Name,
		Count:   1,
		Time:    time.Now().String(),
		Creator: name,
	}
	team.Save()

	return nil
}

// UpdateTeam ... 更新团队
func (s *MemberService) UpdateTeam(ctx context.Context, req *pb.UpdateTeamRequest, resp *pb.Response) error {

	err := dao.UpdateTeamName(req.Id, req.Name)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTeam ... 删除团队
func (s *MemberService) DeleteTeam(ctx context.Context, req *pb.DeleteTeamRequest, resp *pb.Response) error {

	err := dao.DeleteTeam(req.Id)
	if err != nil {
		return err
	}

	return nil
}
