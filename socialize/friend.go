package socialize

import (
	"Douyin-Youth-training/common"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func FriendList(ctx *app.RequestContext, user common.User) {
	ctx.JSON(http.StatusOK, common.UserListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		UserList: user.FriendList,
	})
}
