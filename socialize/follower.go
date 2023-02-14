package socialize

import (
	"Douyin-Youth-training/controller"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func FollowerList(ctx context.Context, c *app.RequestContext) {
	// 解析user id，TODO：正确性待测试
	userId := c.Query("user_id")
	// 根据userid从数据库中读取该id的粉丝列表
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var user controller.User
	db.First(&user, userId)
	// 将要返回的请求体序列化为json格式
	c.JSON(http.StatusOK, controller.UserListResponse{
		Response: controller.Response{
			StatusCode: 0,
		},
		UserList: user.FollowerList,
	})
}
