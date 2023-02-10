package main

import (
	"Douyin-Youth-training/socialize"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func initrouter(h *server.Hertz) {
	apiRouter := h.Group("/douyin")

	// socialize
	apiRouter.GET("/relation/follower/list/", socialize.FollowerList)
	apiRouter.GET("/relation/friend/list/", socialize.FriendList)
}
