package user

import (
	"context"
	"strconv"

	. "workbench-gateway/handler"
	"workbench-gateway/service"
	"workbench-gateway/util"
	pb "workbench-user/proto"
	"workbench/log"
	"workbench/pkg/errno"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// List ... 获取 userList
// @Summary get user_list api
// @Description 通过 group 和 team 获取 user_list
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param last_id query int false "last_id"
// @Param Authorization header string true "token 用户令牌"
// @Param group_id path int true "group_id"
// @Param team_id path int true "team_id"
// @Success 200 {object} ListResponse
// @Router /user/list/{group_id}/{team_id} [get]
func List(c *gin.Context) {
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 limit
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendError(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取 group 和 team
	groupID, err := strconv.Atoi(c.Param("group_id"))
	if err != nil {
		SendError(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}
	teamID, err := strconv.Atoi(c.Param("team_id"))
	if err != nil {
		SendError(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取 page
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendError(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	lastId, err := strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendError(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 构造请求给 list
	listReq := &pb.ListRequest{
		LastId: uint32(lastId),
		Offset: uint32(page * limit),
		Limit:  uint32(limit),
		Team:   uint32(teamID),
		Group:  uint32(groupID),
	}

	// 发送请求
	listResp, err := service.UserClient.List(context.TODO(), listReq)
	if err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendResponse(c, nil, listResp)
}
