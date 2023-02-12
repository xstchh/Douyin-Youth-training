package controller

import (
	"Douyin-Youth-training/common"
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"
)

type VideoListResponse struct {
	common.Response
	VideoList []Video `json:"video_list"`
}

// Publish 检查 token 是否有效，如果有效则将上传的文件保存到public目录下，并返回成功信息
func Publish(ctx context.Context, c *app.RequestContext) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList 返回所有用户的发布的视频列表，在这个例子中使用的是一个示例视频列表
func PublishList(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
