package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义了两个处理视频收藏夹的函数

// FavoriteAction没有实际效果，只是检查token是否有效

/*
	更具体地：
	第一个函数 FavoriteAction 使用*gin.Context作为参数，
	它是一个指向包含当前HTTP请求及其上下文信息的结构体的指针。
	该函数从请求中提取一个 token 查询参数，并检查它是否存在于 usersLoginInfo 映射中。
	如果是，该函数将发送一个状态码为0的JSON响应，表示成功。
	如果 token 不存在，该函数将发送一个带有状态代码1和状态消息“用户不存在”的JSON响应。
*/
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList所有用户都有相同的收藏视频列表
/*
	第二个功能 Favoritelist 它发送一个包含 VideoListResponse 结构体的JSON响应。
	该结构体包含一个状态码为0的 Response 结构体和一个 DemoVideos 的 Video 结构体数组。
	这个函数返回所有用户收藏的视频列表。
*/

func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
