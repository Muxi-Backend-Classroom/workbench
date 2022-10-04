package service

import (
	"context"
	"workbench-team/dao"
	"workbench-team/pkg"
	team "workbench-team/proto"
	"workbench/pkg/errno"
)

// 返回邀请链接 url
func (m *MemberService) GetApplyLink(ctx context.Context, req *team.GetApplyLinkRequest, resp *team.GetApplyLinkResponse) error {
	urlSuffix, err := pkg.GenerateUrl(req.UserId, req.GroupId)
	if err != nil {
		errno.ServerErr(errno.ErrGenerate, err.Error())
	}
	resp.Link = "http://work.muxi-tech.xyz/team/invite/?hash_id=" + urlSuffix
	return nil
}

// 确认邀请通过
func ConfirmApply(ctx context.Context, req *team.ConfirmApplyRequest, resp *team.Response) error {
	infor, err := pkg.ParseUrl(req.Url) // 生成链接的发起人
	if err != nil {
		return errno.ServerErr(errno.ErrParse, err.Error())
	}
	
	group, err := dao.GetGroup(infor.Id)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	
	team, err := dao.GetTeam(infor.Id)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	
	if err = dao.UpdateTeamAndGroup(req.Id, group.Id, team.Id); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	
	apply, err := dao.GetApply(req.Id)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	
	if err = dao.DeleteApply(apply); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	
	return nil
}

// 获取所有apply
func GetApplys(ctx context.Context, req *team.Request, resp *team.ApplysResponse) error {
	rets, err := dao.GetAllApplys()
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	
	resp.Applys = make([]*team.Apply, len(rets))
	copy(resp.Applys, rets)
	return nil
}

func Apply(ctx context.Context, req *team.ApplyRequest, resp *team.Response) error {
	apply := dao.ApplysModel{
		UserId: req.UserId,
	}
	if err := apply.Create(); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
