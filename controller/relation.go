package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RelationAction 检查token是否有效，如果存在就返回一个状态码为0的响应，
// 否则返回状态码为1并且状态信息为 "User doesn't exist" 的响应
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList 返回所有用户的关注列表，数据都是固定的，即只有一个 DemoUser
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}

// FollowerList 返回所有用户的粉丝列表，数据都是固定的，即只有一个 DemoUser
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}

// FriendList 返回所有用户的好友列表，同样，数据都是固定的，即只有一个 DemoUser
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}
