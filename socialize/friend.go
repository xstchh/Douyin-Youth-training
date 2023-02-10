package socialize

import (
	"Douyin-Youth-training/common"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func FriendList(c context.Context, ctx *app.RequestContext) {
	// 解析user id，TODO：正确性待测试
	userId := c.Value("user_id")
	// 根据userid从数据库中读取该id的好友列表
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var user common.User
	db.First(&user, userId)
	// 将要返回的请求体序列化为json格式
	ctx.JSON(http.StatusOK, common.UserListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		UserList: user.FriendList,
	})
}
