package base

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
	// VideoList字段是一个视频列表，NextTime字段是下一次请求的时间戳
}

// Feed 对于每个请求都提供相同的演示视频列表
func Feed(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
