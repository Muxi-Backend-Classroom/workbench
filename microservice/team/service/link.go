package service

import (
	"context"
	"github.com/Muxi-Backend-Classroom/workbench/pkg/errno"
	"workbench-team/pkg"
	team "workbench-team/proto"
)

func (m *MemberService) GetApplyLink(ctx context.Context, req *team.GetApplyLinkRequest, resp *team.GetApplyLinkResponse) error {
	urlSuffix, err := pkg.GenerateUrl(req.UserId)
	if err != nil {
		errno.ServerErr(&errno.Errno{
			Code:    10005,
			Message: "Error occurred while generate url",
		}, err.Error())
	}
	resp.Link = "" + urlSuffix
	return nil
}

func Apply() error {

}

func ConfirmApply() error {

}

func GetApplys() error {

}
