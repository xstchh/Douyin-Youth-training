package socialize

import (
	"Douyin-Youth-training/common"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func FollowerList(ctx *app.RequestContext, user common.User) {
	// 将要返回的请求体序列化为json格式
	ctx.JSON(http.StatusOK, common.UserListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		UserList: user.FollowerList,
	})
}
