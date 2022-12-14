package user

import (
	"context"

	. "workbench-gateway/handler"
	"workbench-gateway/service"
	"workbench-gateway/util"
	pb "workbench-user/proto"
	"workbench/log"
	"workbench/pkg/errno"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetInfo ... 获取 userInfo
func GetInfo(c *gin.Context) {
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取 Ids
	var req GetInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendError(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	if len(req.Ids) == 0 {
		SendResponse(c, nil, &GetInfoResponse{})
		return
	}

	// 构造请求给 getInfo
	var getInfoReq = &pb.GetInfoRequest{}
	for _, id := range req.Ids {
		getInfoReq.Ids = append(getInfoReq.Ids, id)
	}

	// 发送请求
	getInfoResp, err := service.UserClient.GetInfo(context.TODO(), getInfoReq)
	if err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	// 构造返回 response
	var resp GetInfoResponse
	for _, item := range getInfoResp.List {
		resp.List = append(resp.List, UserInfo{
			Id:        item.Id,
			Name:      item.Name,
			RealName:  item.RealName,
			AvatarURL: item.AvatarUrl,
			Email:     item.Email,
		})
	}

	SendResponse(c, nil, resp)
}
