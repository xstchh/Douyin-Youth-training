package socialize

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type User struct {
	Id            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
	FollowerList  []User // TODO:新添加粉丝列表属性，是否最优做法待商榷
}

func FollowerList(ctx *app.RequestContext, user User) {
	// 将要返回的请求体序列化为json格式
	ctx.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: user.FollowerList,
	})
}
