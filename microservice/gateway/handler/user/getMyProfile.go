package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	. "workbench-gateway/handler"
	"workbench-gateway/util"
	"workbench/log"
)

// GetMyProfile ... 获取 myProfile
// @Summary get my_profile api
// @Description 获取 my 完整 user 信息
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} UserProfile
// @Router /user/myprofile [get]
func GetMyProfile(c *gin.Context) {
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	userId := c.MustGet("userId").(uint32)

	user, err := GetUserProfile(userId)

	if err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendResponse(c, nil, user)
}
